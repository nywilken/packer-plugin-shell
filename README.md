### Packer Shell Provisioner Plugin

### Plugin Setup
 - bootstrap a directory with go files
 - have a main.go template; possibly have an auto-generate command.
 - split from packer core with history intact.
 - reap benefits.


### Split process:

 Single plugins with a single type (e.g shell)
 ```
 # Lets start by reducing the Packer repo down to just the
 git clone git@github.com:hashicorp/packer packer-provisioner-shell
 cd packer-provisioner-shell
 git filter-branch --prune-empty --tag-name-filter cat --subdirectory-filter provisioner/shell -- --all
 # Validate file tree

 git clone -depth 1 git@github.com/hashicorp/packer-plugin-scaffolding packer-plugin-shell

 cd packer-plugin-shell
 git remote add shell-provisioner ../.packer-provisioner-shell
 git merge -s ours --no-commit --allow-unrelated-histories shell-provisioner/master
 git read-tree --prefix=shell -u shell-provisioner/master

 ## Fix go mods and imports; packer should be in packer-plugin-sdk
 ## update tests
 go mod tidy
 go get ./...
 go build
 ```

 Plugin with a multiple types (e.g amazon)



