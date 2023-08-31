# Notas Rápidas CLI

Notas Rápidas CLI é uma ferramenta de linha de comando (CLI) minimalista para capturar e gerenciar notas curtas de maneira eficaz. Permite aos usuários adicionar, listar, copiar para a área de transferência e excluir notas com facilidade.

## Recursos

- Adição rápida de notas
- Listagem e exclusão de notas
- Copiar notas para área de transferência
- Limpar todas as notas
- Listagem de notas por período

## Uso

1. Certifique-se de ter Go instalado em seu sistema.
2. Clone este repositório.
3. Navegue para o diretório do repositório.
4. Execute o comando `go build` para compilar o executável.
5. Mova o executável gerado para um local no seu `$PATH`.

### Comandos Disponíveis

- **Adicionar Nota:** `notes add "Conteúdo da Nota"`
- **Listar Notas:** `notes ls`
- **Copiar Nota:** `notes cp <número>`
- **Excluir Nota:** `notes rm <número>`
- **Limpar Notas:** `notes clear`
- **Listar Notas de Ontem:** `notes ls y`
- **Listar Notas da Semana:** `notes ls w`

## Contribuição

Contribuições são bem-vindas! Se você deseja melhorar este projeto, sinta-se à vontade para abrir problemas e enviar pull requests.

## Licença

Este projeto é licenciado sob a [Licença Pública Geral GNU versão 3 (GNU GPLv3)](LICENSE). Consulte o arquivo [LICENSE](LICENSE) para obter detalhes.

---