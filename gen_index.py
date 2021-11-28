import os


def write_sidebar(names):

    lines = ['<!-- docs/_sidebar.md --> ', '', '* [Home](/)']
    #* [Go](./go.md)
    for n in names:
        t = n.split('.')[0].title()
        lines.append(f'* [{t}](./{n})')

    sidebar = './docs/_sidebar.md'
    os.remove(sidebar)
    with open(sidebar, 'w') as f:
        f.writelines('\n'.join(lines))
    
def write_readme(names):
    lines = ['# 二驴日志', '', '', '']
    #* [Go](./go.md)
    for n in names:
        n = n.split('.')[0]
        t = n.title()
        lines.append(f'[{t}](./{n})')
        lines.append('')
    
    
    sidebar = './docs/README.md'
    os.remove(sidebar)
    with open(sidebar, 'w') as f:
        f.writelines('\n'.join(lines))

def get_filesname():
    names = []
    files = os.listdir('docs')
    for file in  files:
        if not file.startswith('_') and not file.endswith('Dockerfile.'):
            names.append(file)
    return names    



if __name__ == '__main__':
    names = get_filesname()
    write_sidebar(names)
    write_readme(names)
