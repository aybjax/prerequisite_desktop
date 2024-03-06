export type NodeType = 'chapter'|'unit'|'topic'|'microtopic'
export type ElementType = (cytoscape.NodeDataDefinition & {type: NodeType}) | (cytoscape.EdgeDataDefinition)