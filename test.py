import pystache
import yaml
import html

with open('template-build.sh','r+') as f:
    template = f.read()
    with open('example.travis.yml') as y:
        obj = yaml.load(y)
        print(html.unescape(pystache.render(template, obj["env"])))
