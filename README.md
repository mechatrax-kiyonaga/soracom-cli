[![go report status](https://goreportcard.com/badge/github.com/soracom/soracom-cli)](https://goreportcard.com/report/github.com/soracom/soracom-cli)
![build-artifacts](https://github.com/soracom/soracom-cli/actions/workflows/build-artifacts.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/soracom/soracom-cli.svg)](https://pkg.go.dev/github.com/soracom/soracom-cli)

# soracom-cli

A command line tool `soracom` to invoke SORACOM API.

# Feature

The `soracom` command:

- supports new APIs on-time. The binary file of the soracom command is automatically generated from the API definition file.

- just works by copying the cross-compiled binary file into the target environment. There is no need to build an environment or solve dependencies.

- constructs a request based on the specified argument and calls the SORACOM API. Response (JSON) from the API is output directly to standard output.
  - This makes it easier to process the output of the soracom command and pass it to another command

- supports bash completion. Please write the following line in .bashrc etc
  ```
  eval "$(soracom completion bash)"
  ```

  if you are a macOS user, you probably need to either:
  1. use `bash` version >= 4.0, or
  2. use `brew install bash-completion` instead of using Xcode version of bash-completion and then add the following to either your `.bash_profile` or `.profile`:

  ```
  if [ -f $(brew --prefix)/etc/bash_completion ]; then
    . $(brew --prefix)/etc/bash_completion
  fi
  ```
  otherwise you might be getting the error like the following:
  ```
  -bash: __ltrim_colon_completions: command not found
  ```

- supports zsh completion. The generated completion script by running the following command should be put somewhere in your $fpath named `_soracom`
  ```
  soracom completion zsh
  ```

# How to install

## Using Mac (macOS) or Linux, installing by homebrew

```shell
brew tap soracom/soracom-cli
brew install soracom-cli
brew install bash-completion
```

## Using Linux, installing with snap

```shell
sudo snap install soracom
```

In this case, you may want to connect the snap to `dot-soracom` interface which is for reading / writing profile files in `$HOME/.soracom`:

```shell
snap connect soracom:dot-soracom
```

You may also want to have the following line in your `.bashrc` etc. to use `$HOME/.soracom` as the profiles directory instead of `$SNAP_USER_DATA/.soracom` (i.e. `$HOME/snap/soracom/<revision>/.soracom`)

```bash
export SORACOM_PROFILE=$HOME/.soracom
```

## In other cases

By running one of the following commands, the latest version of `soracom` command will be installed.

If you have a permission to write a file into `/usr/local/bin` directory (e.g. you are `root` user), please run the command below:

```shell
curl -fsSL https://raw.githubusercontent.com/soracom/soracom-cli/master/install.sh | bash
```

If you do not have a permission to write a file into `/usr/local/bin` directory, please run either of the following commands.

If you are in sudoers and want to install `soracom` command to `/usr/local/bin`:

```shell
curl -fsSL https://raw.githubusercontent.com/soracom/soracom-cli/master/install.sh | sudo bash
```

or

If you are not in sudoers or want to install `soracom` command to other directory e.g. `$HOME/bin`:

```shell
mkdir -p "$HOME/bin"
curl -fsSL https://raw.githubusercontent.com/soracom/soracom-cli/master/install.sh | BINDIR="$HOME/bin" bash
```

You can change `"$HOME/bin"` in the command above to wherever you want.

If you want to upgrade the `soracom` command, you can just run the same command you used to install `soracom` again.

If you want to uninstall the `soracom` command, you can just remove `soracom` executable file you have installed. (You may want to remove `$HOME/.soracom/` directory which contains profiles for the `soracom` command.)

If the commands above did not work well, or if you want to install older version of `soracom` command, you can download a package file that match the environment of the target from [Releases page](https://github.com/soracom/soracom-cli/releases), unpack it, and place the executable file in the directory where included in `PATH`.


# How to use

## Basic usage

First of all, create a profile by running the following command:

```
soracom configure
```

You will be asked which coverage type to use.

```
Please select which coverage type to use.

1. Global
2. Japan

select (1-2) >
```

Please select the coverage type which you mainly use. In most cases, please select Global. If you live in Japan and use SIM cards in Japan, please select Japan.

Next you will be asked about the authentication method.

```
Please select which authentication method to use.

1. Input AuthKeyId and AuthKey * Recommended *
2. Input Operator credentials (Operator Email and Password)
3. Input SAM credentials (OperatorId, User name and Password)

select (1-3) >
```

Please select 1 if AuthKey (authentication key) has been issued to SAM user or root account.
(For details on how to issue an authentication key to SAM users, please see [Using SORACOM Access Management to Manage Operation Access](https://dev.soracom.io/en/start/sam/).

Thereafter, when executing the soracom command, an API call is made using the authentication information entered here.



## Advanced usage

### Use multiple profiles

If you have multiple SORACOM accounts or want to use multiple SAM users differently, specify the --profile option to configure and set the profile name.

```
soracom configure --profile user1
  :
  (Enter information for user1)

soracom configure --profile user2
  :
  (Enter information for user2)
```

This will create profiles named user1 and user2.
To use the profile, specify the --profile option in addition to the normal command.

```
soracom subscribers list --profile user1
  :
  (SIM list for user1 will be displayed)

soracom groups list --profile user2
  :
  (Group list for user2 will be displayed)
```


### Create a profile for API Sandbox

It is possible to use soracom-cli for setting up [SORACOM API Sandbox](https://dev.soracom.io/en/docs/api_sandbox/) environment.

In order to create a profile for sandbox, use `configure-sandbox` subcommand.

```
soracom configure-sandbox
```

By answering to the questions prompted, a profile named `sandbox` will be created. By using the `sandbox` profile, you can issue commands to the API sandbox as follows.

```
soracom subscribers list --profile sandbox
```

You can use commands dedicated for the sandbox.

```
soracom sandbox subscribers create --profile sandbox
```

You can use different profile name.

```
soracom configure-sandbox --profile test
soracom sandbox subscribers create --profile test
```

In order to make it easier to use from shell scripts etc., all the parameters necessary for profile creation can be specified with arguments.

```
soracom configure-sandbox --coverage-type jp --auth-key-id="$AUTHKEY_ID" --auth-key="$AUTHKEY" --email="$EMAIL" --password="$PASSWORD"
```


### Call API via proxy

Set `http://your-proxy-name:port` to HTTP_PROXY environment variable, then execute soracom command.

e.g.) For Linux / Mac:
Assume that the address of the proxy server is 10.0.1.2 and the port number is 8080
```
export HTTP_PROXY=http://10.0.1.2:8080
soracom subscribers list
```

Or

```
HTTP_PROXY=http://10.0.1.2:8080 soracom subscribers list
```


### Trouble shooting

If you get an error message like the following:

```
Error: Permissions for the file 'path/to/default.json' which contains your credentials are too open.
It is required that your credential files are NOT accessible by others.
```

Please try the following to fix it:

```
soracom unconfigure
soracom configure
```

i.e. perform `unconfigure` and then `configure` again in order to re-create a credentials file with appropriate permissions.


# How to build / test

For developers who want to build from source or for those who wish to make a pull request such as bug fix / function addition, please build and test in one of the following ways.

## How to build in a local environment (Linux / Mac OS X)

In the environment where Go is installed, run the build script as follows:

```
VERSION=1.2.3
./scripts/build.sh $VERSION
```

Here 1.2.3 is the version number. Please specify an appropriate number.

If the build succeeds, then run the test:

```
export SORACOM_AUTHKEY_ID_FOR_TEST=...   # set AuthKey ID & AuthKey of a Soracom operator (account) to use the API sandbox.
export SORACOM_AUTHKEY_FOR_TEST=...
./test/test.sh $VERSION
```


# How to release

```
VERSION=1.2.3                         # => specify a version number to be released
./scripts/build.sh $VERSION           # => build a version to be released
./test/test.sh $VERSION               # => test the version
# commit & push all changes to github
./scripts/release.sh $VERSION         # => release the version to GitHub
# edit the release on github.com release page
./scripts/update-homebrew-formula.sh $VERSION $GITHUB_USERNAME $GITHUB_EMAIL
./scripts/build-snap.sh $VERSION
./scripts/release-snap.sh $VERSION
./scripts/build-lambda-layer.sh $VERSION
./scripts/release-lambda-layer.sh $VERSION $AWS_PROFILE   # => this command releases the layer to all regions (except ap-east-1)
```
