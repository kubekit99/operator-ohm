// Copyright 2018 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build v2

package helmv2

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	// "gopkg.in/src-d/go-git.v4/storage/memory"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	"k8s.io/helm/pkg/chartutil"
	cpb "k8s.io/helm/pkg/proto/hapi/chart"
)

type source struct {
	chartLocation *av1.ArmadaChartSource
}

// Downloads the chart local in case the Chart has not been bundled with the operator
func (m source) getChart() (*cpb.Chart, error) {
	var pathToChart string
	var err error
	switch m.chartLocation.Type {
	case "git":
		pathToChart, err = m.gitClone()
	case "tar":
		pathToChart, err = m.getTarball()
	case "local":
		pathToChart = m.chartLocation.Location
	}

	if err != nil {
		return nil, err
	}

	chart, err := chartutil.LoadDir(pathToChart)
	if err != nil {
		return nil, err
	}

	// err = m.sourceCleanup(pathToChart)
	// if err != nil {
	// 	// TODO: log a warning
	// }

	return chart, nil
}

// '''Clone a git repository from ``repo_url`` using the reference ``ref``.
// :param repo_url: URL of git repo to clone.
// :param ref: branch, commit or reference in the repo to clone. Default is
//     'master'.
// :param proxy_server: optional, HTTP proxy to use while cloning the repo.
// :param auth_method: Method to use for authenticating against the repository
//     specified in ``repo_url``.  If value is "SSH" Armada attempts to
//     authenticate against the repository using the SSH key specified under
//     ``CONF.ssh_key_path``. If value is None, authentication is skipped.
//     Valid values include "SSH" or None. Note that the values are not
//     case sensitive. Default is None.
// :returns: Path to the cloned repo.
// :raises GitException: If ``repo_url`` is invalid or could not be found.
// :raises GitAuthException: If authentication with the Git repository failed.
// :raises GitProxyException: If the repo could not be cloned due to a proxy
//     issue.
// :raises GitSSHException: If the SSH key specified by ``CONF.ssh_key_path``
//     could not be found and ``auth_method`` is "SSH".
// '''
func (m *source) gitClone() (string, error) {

	log.Info("Cloning :", "location", m.chartLocation.Location)

	var repoUser string
	repoURL, err := url.Parse(m.chartLocation.Location)
	if err != nil {
		return "", err
	} else {
		repoUser = repoURL.User.Username()
	}

	// normalizedURL := repoURL.RawPath
	normalizedURL := m.chartLocation.Location
	log.Info("Cloning :", "url", normalizedURL)

	// TODO(jeb): AuthMethod and SSH is still WIP
	var auth transport.AuthMethod
	if m.chartLocation.AuthMethod == "ssh" {
		sshPrivateKey := "/home/" + repoUser + "/.ssh/id_rsa"
		insecureIgnoreHostKey := true
		signer, err := ssh.ParsePrivateKey([]byte(sshPrivateKey))
		if err != nil {
			return "", err
		}
		sshauth := &gitssh.PublicKeys{User: repoUser, Signer: signer}
		if insecureIgnoreHostKey {
			sshauth.HostKeyCallback = ssh.InsecureIgnoreHostKey()
		}
		auth = sshauth

		return "", errors.New("ssh method not supported")

	} else if m.chartLocation.AuthMethod == "http" {
		username := ""
		password := ""
		httpauth := &githttp.BasicAuth{Username: username, Password: password}

		auth = httpauth

		return "", errors.New("http method with authmethod not supported")
	}

	tempDir, err := ioutil.TempDir("/", "armada")
	if err != nil {
		return "", err
	}

	// proxy_server := m.chartLocation.ProxyServer
	ref_spec := m.chartLocation.Reference

	repo, err := m.goGitClone(normalizedURL, tempDir, auth)
	if err != nil {
		return tempDir, err
	}
	// err = m.goGitFetch("", repo, auth, ref_spec)
	// if err != nil {
	// 	return tempDir, err
	// }
	// err = m.goGitCheckout("", repo, "FETCH_HEAD", "")
	err = m.goGitCheckout("", repo, "", ref_spec)
	if err != nil {
		return tempDir, err
	}

	return tempDir + "/" + m.chartLocation.Subpath, err
}

// Downloads the char tarball from the URL.
func (m *source) getTarball() (string, error) {
	tarballPath, err := m.downloadTarball(false)
	if err != nil {
		return "", err
	}
	return m.extractTarball(tarballPath)
}

// downloadTarball Downloads a tarball to /tmp and returns the path
func (m *source) downloadTarball(verify bool) (string, error) {
	file, err := ioutil.TempFile("/", "armada")
	if err != nil {
		return "", err
	}
	response, err := http.Get(m.chartLocation.Location)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	file.Write(body)

	return file.Name(), nil
}

// extractTarball extracts a tarball to /tmp and returns the path
func (m *source) extractTarball(tarballPath string) (string, error) {
	if _, err := os.Stat(tarballPath); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("%s does not exist", tarballPath)
		}
		return "", err
	}

	tempDir, err := ioutil.TempDir("/", "armada")
	if err != nil {
		return "", err
	}

	fileContents, err := os.Open(tarballPath)
	if err != nil {
		return "", err
	}

	gzr, err := gzip.NewReader(fileContents)
	if err != nil {
		return "", err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	done := false
	for !done {
		if err := m.readFromArchive(tr, tempDir); err != nil {
			if err != io.EOF {
				return "", err
			}
			// io.EOF means there's no more data to be read
			done = true
		}
	}
	return tempDir, nil
}

// readFromArchive reads a an item from tr, saves it to dir, then move tr to the next item
func (m *source) readFromArchive(tr *tar.Reader, dir string) error {
	header, err := tr.Next()
	if err != nil {
		// This catches EOF, which means that we're done
		return err
	}

	if header == nil {
		// if the header is nil, just skip it (not sure how this happens)
		return nil
	}

	target := filepath.Join(dir, header.Name)

	switch header.Typeflag {
	case tar.TypeDir:
		if _, err := os.Stat(target); err != nil {
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		}
	case tar.TypeReg:
		f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
		if err != nil {
			return err
		}

		// copy over contents
		if _, err := io.Copy(f, tr); err != nil {
			return err
		}

		// manually close here after each file operation; defering would cause each file close
		// to wait until all operations have completed.
		f.Close()
	}
	return nil
}

func (m *source) sourceCleanup(chart_path string) error {
	// TODO(Ian): Finish this method
	if _, err := os.Stat(chart_path); err == nil {
		err := os.RemoveAll(chart_path)
		if err != nil {
			log.Info("Could not delete the path %s", chart_path)
			return err
		}
	} else {
		log.Info("Could not find the chart path %s to delete.", chart_path)
	}

	return nil
}

// Init initializes a local git repository and sets the remote origin
func (m *source) goGitInit(repoURL string, root string) (*git.Repository, error) {
	log.Info("goGitInit :", "url", repoURL)

	existing, err := git.PlainOpen(root)
	if err == nil {
		return existing, nil
	}
	if err != git.ErrRepositoryNotExists {
		return nil, err
	}
	err = os.RemoveAll(root)
	if err != nil {
		return nil, fmt.Errorf("unable to clean repo at %s: %v", root, err)
	}
	err = os.MkdirAll(root, 0755)
	if err != nil {
		return nil, err
	}
	repo, err := git.PlainInit(root, false)
	if err != nil {
		return nil, err
	}
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: git.DefaultRemoteName,
		URLs: []string{repoURL},
	})
	return repo, err
}

// goGitClone the remote using go-git
func (m *source) goGitClone(repoURL string, root string, auth transport.AuthMethod) (*git.Repository, error) {
	log.Info("goGitClone :", "url", repoURL)

	existing, err := git.PlainOpen(root)
	if err == nil {
		return existing, nil
	}
	if err != git.ErrRepositoryNotExists {
		return nil, err
	}
	err = os.RemoveAll(root)
	if err != nil {
		return nil, fmt.Errorf("unable to clean repo at %s: %v", root, err)
	}
	err = os.MkdirAll(root, 0755)
	if err != nil {
		return nil, err
	}

	options := &git.CloneOptions{URL: repoURL}
	if auth != nil {
		options.Auth = auth
	}
	repo, err := git.PlainClone(root, false, options)
	if err != nil {
		return nil, err
	}

	return repo, err
}

// goGitFetch fetches the remote using go-git
func (m *source) goGitFetch(root string, repo *git.Repository, auth transport.AuthMethod, ref_spec string) error {
	log.Info("goGitFetch :", "root", root)

	if repo == nil {
		if root == "" {
			return errors.New("root and repo are nil")
		}

		existing, err := git.PlainOpen(root)
		if err != nil {
			return err
		}
		repo = existing
	}

	options := &git.FetchOptions{
		RemoteName: git.DefaultRemoteName,
		Tags:       git.AllTags,
		Force:      false,
	}
	if auth != nil {
		options.Auth = auth
	}
	if ref_spec != "" {
		//JEB: can't get the equivalent of git fetch origin hash to work yet
		refSpec := config.RefSpec(plumbing.NewHashReference("FETCH_HEAD", plumbing.NewHash(ref_spec)).String())
		options.RefSpecs = make([]config.RefSpec, 0)
		options.RefSpecs = append(options.RefSpecs, refSpec)
	}
	err := repo.Fetch(options)
	if err == git.NoErrAlreadyUpToDate {
		return nil
	}
	return err
}

// goGitCheckout fetches the remote using go-git
func (m *source) goGitCheckout(root string, repo *git.Repository, branch_name string, hash_value string) error {
	log.Info("goGitCheckout :", "root", root)

	if repo == nil {
		if root == "" {
			return errors.New("root and repo are nil")
		}

		existing, err := git.PlainOpen(root)
		if err != nil {
			return err
		}
		repo = existing
	}

	workTree, err := repo.Worktree()
	if err != nil {
		return err
	}

	options := &git.CheckoutOptions{}
	if branch_name != "" {
		options.Branch = plumbing.ReferenceName(branch_name)
	}
	if hash_value != "" {
		options.Hash = plumbing.NewHash(hash_value)
	}

	err = workTree.Checkout(options)
	if err == git.NoErrAlreadyUpToDate {
		return nil
	}
	return err
}
