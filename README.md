<table style="border-collapse: collapse; border: none">
  <tr>
    <td style="border: none" width="256">
      <img src="./logo.png" alt="Cyphernetes Logo" width="256">
    </td>
    <td style="border: none; padding-left: 20px">
      <h1>Cyphernetes</h1>
      <p>Cyphernetes is a command-line interface (CLI) tool designed to manage Kubernetes resources using a query language inspired by Cypher, the query language of Neo4j. It provides a more intuitive way to interact with Kubernetes clusters, allowing users to express complex operations as graph-like queries.</p>
    </td>
  </tr>
</table>

## Why Cyphernetes?

Kubernetes management often involves dealing with complex and verbose command-line instructions. Cyphernetes simplifies this complexity by introducing a declarative query language that can express these instructions in a more readable and concise form. By leveraging a query language similar to Cypher, users can efficiently perform CRUD operations on Kubernetes resources, visualize resource connections, and manage their Kubernetes clusters with greater ease and flexibility.

## Usage

Cyphernetes offers two main commands: `quote` and `shell`. Below are examples of how to use these commands with different types of queries.

### Quote Command

The `quote` command is used for running single Cyphernetes queries from the command line. 

Example usage:

```bash
$ cyphernetes quote "MATCH (d:Deployment {name: 'nginx'}) RETURN d"
```

This command retrieves information about a Deployment named 'nginx'.

### Shell Command
The shell command launches an interactive shell where you can execute multiple queries in a session.

To start the shell:

```bash
$ cyphernetes shell
```
Within the shell, you can run queries interactively. Here are some examples:

### Basic Node Match

```graphql
MATCH (d:Deployment) RETURN d
```
This query lists all Deployment resources.

### Node with Properties

```graphql
MATCH (d:Deployment {app: 'nginx'}) RETURN d
```
Retrieves Deployments where the app label is 'nginx'.

### Multiple Nodes

```graphql
# With a relationship
MATCH (d:Deployment)->(s:Service) RETURN d, s

# Multiple matches
MATCH (d:Deployment), (s:Service) RETURN d, s
```
Lists Deployments and their associated Services.

### Node with Multiple Properties

```graphql
MATCH (s:Service {type: 'LoadBalancer', region: 'us-west'}) RETURN s
```
Finds Services of type 'LoadBalancer' in the 'us-west' region.

### Returning specific properties

```graphql
MATCH (s:Service) RETURN s[*].status.LoadBalancer
```

### Multiple matches and returns

```graphql
MATCH (d:Deployment)->(s:Service) RETURN d[*].metadata.name, s[*].status.LoadBalancer
```

Remember to type exit or press Ctrl-C to leave the shell.

## Development

Cyphernetes is written in Go and utilizes a parser generated by goyacc to interpret the custom query language.

### Prerequisites

- Go (1.16 or later)
- goyacc (for generating the parser)
- Make (for running make commands)

### Getting Started

To get started with development:

1. Clone the repository:
    ```bash
    git clone https://github.com/avitaltamir/cyphernetes.git
    ```

2. Navigate to the project directory:
    ```bash
    cd cyphernetes
    ```

### Building the Project

Use the Makefile commands to build the project:

- Test & Build:
    ```bash
    $ make
    ```
- To build the binary:
    ```bash
    make build
    ```

- To run tests:
    ```bash
    $ make test
    ```

- To generate the grammar parser:
    ```bash
    $ make gen-parser
    ```

- To clean up the build:
    ```bash
    $ make clean
    ```

### Contributing

Contributions are welcome! Please feel free to submit pull requests, open issues, and provide feedback.

## Project Roadmap

The project is at it's earliest milestone and supports performing GET operations.
The Cypher-like grammer implementation is incomplete, still missing:
* CREATE, SET, DELETE clauses
* Match clauses with more than 2 comma-separated Node patterns
* Match clauses with multiple relationship patterns
* Relationships between more than 2 nodes
* Relationships pattern tokens (relationship arrows are currently similar to commas)
  This will be required later when we introduce more complex K8s operations


### Initial Project Setup

- [x] Initialize the project repository.
- [x] Set up version control with Git.
- [x] Create and document the project directory structure.
- [x] Choose a Go package management tool and initialize the package.
- [x] Set up a Go workspace with the necessary Go modules.

### Tooling and Framework

- [x] Set up a testing framework using Go's built-in testing package.
- [x] Configure a continuous integration service.
- [x] Establish linting and code formatting tools.
- [x] Implement logging and debug output mechanisms.

### Lexer and Parser Development

- [x] Create the basic lexer with support for initial tokens.
- [x] Develop a yacc file for the initial grammar rules.
- [x] Write unit tests for basic tokenization.
- [x] Implement a basic parser to handle `MATCH` queries.
- [x] Test and debug the lexer and parser with simple queries.

### Kubernetes Client Integration

- [x] Evaluate and select a Go Kubernetes client library.
- [x] Set up authentication and configuration for accessing a Kubernetes cluster.
- [x] Implement a wrapper around the Kubernetes client to execute basic queries.
- [x] Develop mapping logic to convert parsed queries into Kubernetes API calls.
- [ ] Test Kubernetes client integration with mock and real clusters.
- [ ] Add support for complex queries involving multiple Kubernetes resources.

### Expanding Lexer and Parser

- [x] Add support for additional tokens (e.g., braces, commas, relationship types).
- [x] Extend grammar rules to cover node properties and relationships.
- [ ] Implement parsing logic for `CREATE`, `SET`, and `DELETE` keywords.
- [ ] Refine error handling for syntax and parsing errors.
- [ ] Optimize lexer and parser for performance.

### Interactive Shell Interface

- [x] Basic shell interface for inputting queries and displaying results.
- [ ] Add help and documentation to the shell.
- [ ] Autocompletion.
- [ ] Syntax highlighting.
- [ ] Test shell with various input scenarios.

## Cypher-Like Query Language Parser Roadmap

The goal of this roadmap is to incrementally develop a parser that can handle a Cypher-like query language. The final version should support complex queries involving `MATCH`, `RETURN`, `CREATE`, `SET`, and `DELETE` statements.

### Phase 1: Basic MATCH Support

- [x] Support for basic `MATCH` queries (e.g., `MATCH (k:Kind)`).
- [x] Write unit tests for basic `MATCH` query parsing.

### Phase 2: RETURN Clause

- [x] Implement parsing of the `RETURN` clause.
- [x] Update the lexer to recognize the `RETURN` keyword.
- [x] Extend the yacc grammar to include `RETURN` statement rules.
- [x] Write unit tests for queries with `RETURN` clauses.

### Phase 3: Node Properties

- [x] Extend the parser to handle node properties.
- [x] Update the lexer to recognize curly braces and commas.
- [x] Update the yacc file to handle node properties syntax.
- [x] Write unit tests for `MATCH` queries with node properties.

### Phase 4: Relationships

- [x] Support parsing of relationships in `MATCH` queries.
- [x] Extend the yacc grammar to handle relationship patterns.
- [x] Write unit tests for `MATCH` queries involving relationships.
- [ ] Update the lexer to recognize relationship pattern tokens (e.g., `-[]->`).
- [ ] Support relationships between more than 2 nodes.

### Phase 5: Advanced MATCH Support
- [ ] Match Clauses to contain NodePatternLists instead of a single tuple of Node/ConnectedNode
- [ ] Support more than 2 comma-separated NodePatternLists.

### Phase 6: CREATE Statement

- [ ] Add support for `CREATE` statements.
- [ ] Update the lexer to recognize the `CREATE` keyword.
- [ ] Extend the yacc grammar to parse `CREATE` statements.
- [ ] Write unit tests for `CREATE` statement parsing.

### Phase 7: SET Clause

- [ ] Implement parsing of the `SET` clause.
- [ ] Update the lexer to recognize the `SET` keyword and property assignment syntax.
- [ ] Extend the yacc grammar to include `SET` statement rules.
- [ ] Write unit tests for queries with `SET` clauses.

### Phase 8: DELETE Statement

- [ ] Add support for `DELETE` statements.
- [ ] Update the lexer to recognize the `DELETE` keyword.
- [ ] Extend the yacc grammar to parse `DELETE` statements.
- [ ] Write unit tests for `DELETE` statement parsing.

### Phase 9: Complex Query Parsing

- [ ] Combine all elements to support full query parsing.
- [ ] Ensure the lexer and yacc grammar can handle complex queries with multiple clauses.
- [ ] Write unit tests for parsing full queries including `MATCH`, `RETURN`, `CREATE`, `SET`, and `DELETE`.

## License

Cyphernetes is open-sourced under the MIT license. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the [Neo4j](https://neo4j.com/) community for the inspiration behind the query language.

## Authors

- _Initial work_ - [Avital Tamir](https://github.com/avitaltamir)
