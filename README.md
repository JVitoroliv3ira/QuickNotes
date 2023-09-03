# Quick Notes - Gerenciador de Notas Rápidas

Bem-vindo ao QuickNotes! Esta ferramenta de linha de comando minimalista foi criada para facilitar a captura e o gerenciamento de notas curtas de maneira eficaz. Com recursos simples, mas poderosos, você pode adicionar, listar, copiar para a área de transferência e excluir notas com facilidade.

## Recursos

- **Adição Rápida de Notas**: Capture suas ideias instantaneamente com o comando `notes create "Insira aqui o conteúdo da nota"`.
- **Listagem de Notas**: Visualize todas as suas notas existentes usando o comando `notes ls`.
- **Cópia para Área de Transferência**: Copie o conteúdo de uma nota para a área de transferência com o comando `notes cp <número>`.
- **Exclusão de Notas**: Remova uma nota específica digitando `notes rm <número>`.
- **Limpeza de Todas as Notas**: Libere espaço e mantenha sua lista de notas organizada com o comando `notes clear`.

## Como Começar

### Opção 1: Clone e Compile

1. **Clone o Repositório**: Clone este repositório para sua máquina utilizando o comando:

   ```
   git clone https://github.com/JVitoroliv3ira/QuickNotes
   ```

2. **Navegue até o Diretório do Repositório**:

   ```
   cd QuickNotes
   ```

3. **Compile o Executável**:

   ```
   go build -o notes main.go
   ```

4. **Mova o Executável para o Binário Local**:
    - No Linux, utilize o comando `sudo mv notes /usr/local/bin` para adicionar o executável ao seu `$PATH`. Isso permitirá que você acesse o QuickNotes de qualquer lugar no seu sistema.

### Opção 2: Baixe o Executável do Release

1. **Baixe o Executável do Release**: Certifique-se de baixar a versão mais recente do QuickNotes em [releases](https://github.com/JVitoroliv3ira/QuickNotes/releases).

2. **Mova o Executável para o Binário Local**:
    - No Linux, utilize o comando `sudo mv <arquivo> /usr/local/bin` para adicionar o executável ao seu `$PATH`. Isso permitirá que você acesse o QuickNotes de qualquer lugar no seu sistema.

### Comece a Usar

5. **Crie Suas Notas**: Comece a adicionar notas curtas imediatamente com o comando `notes create "Insira aqui o conteúdo da nota"`.

## Comandos Disponíveis

- **Adicionar Nota**: `notes create "Insira aqui o conteúdo da nota"`
- **Listar Notas**: `notes ls`
- **Copiar Nota**: `notes cp <número>`
- **Excluir Nota**: `notes rm <número>`
- **Limpar Notas**: `notes clear`

## Contribuição

Contribuições são bem-vindas! Se você deseja melhorar este projeto, sinta-se à vontade para abrir issues e enviar pull requests. Juntos, podemos tornar o QuickNotes ainda melhor.

## Licença

Este projeto é licenciado sob a [Licença Pública Geral GNU versão 3 (GNU GPLv3)](LICENSE). Consulte o arquivo [LICENSE](LICENSE) para obter detalhes.