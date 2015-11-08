trap 'catch && finally' ERR EXIT
set -e

working_directory={{working_directory}}
repository={{repository}}
branch={{branch}}

mkdir -p "$working_directory"
git clone "$repository" "$working_directory"
cd "$working_directory"
git checkout "$branch"

try(){
    {{#before_install}}
    echo "before_install"
    echo {{.}}
    {{.}}
    {{/before_install}}
    {{#install}}
    echo "install"
    echo {{.}}
    {{.}}
    {{/install}}
    {{#before_script}}
    echo "before_script"
    echo {{.}}
    {{.}}
    {{/before_script}}
    {{#script}}
    echo "script"
    echo {{.}}
    {{.}}
    {{/script}}
    {{#after_success}}
    echo "after_success"
    echo {{.}}
    {{.}}
    {{/after_success}}
    {{#after_script}}
    echo "after_script"
    echo {{.}}
    {{.}}
    {{/after_script}}
    {{#before_deploy}}
    echo "before_deploy"
    echo {{.}}
    {{.}}
    {{/before_deploy}}
    {{#deploy}}
    echo "deploy"
    echo {{.}}
    {{.}}
    {{/deploy}}
    {{#after_deploy}}
    echo "after_deploy"
    echo {{.}}
    {{.}}
    {{/after_deploy}}
}

catch(){
    echo "CAUGHT"
    {{#after_failure}}
    echo "after_failure"
    {{.}}
    {{/after_failure}}
}

finally(){
    echo "FINALLY"
}

try
finally
