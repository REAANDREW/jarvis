trap 'catch && finally' ERR EXIT
set -e

working_directory={{WorkingDirectory}}
repository={{Repository}}
branch={{Branch}}

mkdir -p "$WorkingDirectory"
git clone "$Repository" "$WorkingDirectory"
cd "$WorkingDirectory"
git checkout "$Branch"

try(){
    {{#BeforeInstall}}
echo "BeforeInstall"
    echo {{{.}}}
    {{{.}}}
    {{/BeforeInstall}}
    {{#Install}}
echo "Install"
    echo {{{.}}}
    {{{.}}}
    {{/Install}}
    {{#BeforeScrip}}
echo "BeforeScrip"
    echo {{{.}}}
    {{{.}}}
    {{/BeforeScrip}}
    {{#Script}}
echo "Script"
    echo {{{.}}}
    {{{.}}}
    {{/Script}}
    {{#AfterSuccess}}
echo "AfterSuccess"
    echo {{{.}}}
    {{{.}}}
    {{/AfterSuccess}}
    {{#AfterScript}}
echo "AfterScript"
    echo {{{.}}}
    {{{.}}}
    {{/AfterScript}}
    {{#BeforeDeploy}}
echo "BeforeDeploy"
    echo {{{.}}}
    {{{.}}}
    {{/BeforeDeploy}}
    {{#Deploy}}
echo "Deploy"
    echo {{{.}}}
    {{{.}}}
    {{/Deploy}}
    {{#AfterDeploy}}
echo "AfterDeploy"
    echo {{{.}}}
    {{{.}}}
    {{/AfterDeploy}}
}

catch(){
    echo "CAUGHT"
    {{#AfterFailure}}
echo "AfterFailure"
    {{{.}}}
    {{/AfterFailure}}
}

finally(){
    echo "FINALLY"
}

try
finally
