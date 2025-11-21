# Example Documentation with Mermaid Diagrams

This is an example markdown file demonstrating the annotation strategy for automatic mermaid diagram updates.

## Architecture Overview

The following diagram shows the import relationships in the grapher package:

<!-- mermaid-embed-start:grapher-imports -->
```mermaid
---
---

stateDiagram-v2
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "maps"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "slices"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "fmt"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "iter"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "io"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "errors"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "strings"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "github.com/charmbracelet/log"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "github.com/spf13/pflag"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "bytes"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "github.com/daanv2/go-code-grapher/pkg/extensions/xos"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "github.com/daanv2/go-code-grapher/pkg/markdown"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs" --> "github.com/daanv2/go-code-grapher/pkg/extensions/xslices"
    "github.com/daanv2/go-code-grapher/pkg/grapher/mermaid" --> "strings"
    "github.com/daanv2/go-code-grapher/pkg/grapher/mermaid" --> "errors"
    "github.com/daanv2/go-code-grapher/pkg/grapher/mermaid" --> "github.com/daanv2/go-code-grapher/pkg/grapher/graphs"
    "github.com/daanv2/go-code-grapher/pkg/grapher/mermaid" --> "github.com/daanv2/go-code-grapher/pkg/grapher/graphs/state-diagrams"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs/statediagrams" --> "github.com/daanv2/go-code-grapher/pkg/grapher/graphs"
    "github.com/daanv2/go-code-grapher/pkg/grapher/graphs/statediagrams" --> "github.com/spf13/pflag"
    "github.com/daanv2/go-code-grapher/pkg/grapher" --> "github.com/daanv2/go-code-grapher/pkg/grapher/graphs"
    "github.com/daanv2/go-code-grapher/pkg/grapher" --> "github.com/daanv2/go-code-grapher/pkg/grapher/graphs/state-diagrams"
    "github.com/daanv2/go-code-grapher/pkg/grapher" --> "github.com/daanv2/go-code-grapher/pkg/grapher/mermaid"
    "github.com/daanv2/go-code-grapher/pkg/grapher" --> "fmt"
    "github.com/daanv2/go-code-grapher/pkg/grapher" --> "github.com/daanv2/go-code-grapher/pkg/ast"

```
<!-- mermaid-embed-end:grapher-imports -->

This diagram can be automatically updated by running:

```bash
go-code-grapher imports \
  --markdown-embed-into ./docs/EXAMPLE.md \
  --markdown-embed-id grapher-imports \
  --dir ./pkg/grapher \
  --graph-only false
```

## Another Section

You can have multiple annotated sections in the same file:

<!-- mermaid-embed-start:pkg-structure -->
```mermaid
stateDiagram-v2
    note right of Placeholder: This is another placeholder
    Placeholder --> Data
```
<!-- mermaid-embed-end:pkg-structure -->

Each section has a unique ID and can be updated independently.
