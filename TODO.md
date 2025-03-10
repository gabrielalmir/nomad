# Nomad

Interface Única para Instalação de Pacotes em Diferentes Distros Linux.

## Requisitos Funcionais

### Interface de Uso do Nomad

- Instalar
```bash
(sudo) nomad install <pacote>
```

- Procurar
```bash
(sudo) nomad search <pacote>
```

- Atualizar lista de repositórios
```bash
(sudo) nomad update
```

- Atualizar sistema
```bash
(sudo) nomad upgrade
```

- Remover pacote do sistema
```bash
(sudo) nomad remove <pacote>
```

- [x] Menu de Ajuda
```bash
nomad
```

### Instalação em outros gerenciadores de pacotes

- Instalar
```bash
(sudo) apt install <pacote> [flags opcionais]
```

- Procurar
```bash
(sudo) apt search <pacote>
(sudo) pacman -Q <pacote>
```

- Atualizar lista de repositórios
```bash
(sudo) apt update
(sudo) pacman -Sy
```

- Atualizar sistema
```bash
(sudo) apt upgrade
(sudo) pacman -Syu
```

- Remover pacote do sistema
```bash
(sudo) apt remove <pacote>
(sudo) pacman -R <pacote>
```

- Menu de Ajuda
```bash
usage:  pacman <operation> [...]
operations:
    pacman {-h --help}
    pacman {-V --version}
    pacman {-D --database} <options> <package(s)>
    pacman {-F --files}    [options] [file(s)]
    pacman {-Q --query}    [options] [package(s)]
    pacman {-R --remove}   [options] <package(s)>
    pacman {-S --sync}     [options] [package(s)]
    pacman {-T --deptest}  [options] [package(s)]
    pacman {-U --upgrade}  [options] <file(s)>

use 'pacman {-h --help}' with an operation for available options
```

## Quando precisa de SUDO?

- Deu ruim? Deseja tentar novamente com sudo?
- Pergunta para o sistema operacional se o usuário logado é root (se não for, sugerir utilizar sudo)

Obs: Os argumentos extras, --quiet, --yes, --force, --fix
