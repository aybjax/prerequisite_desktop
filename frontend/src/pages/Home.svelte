<script lang="ts">
    // @ts-ignore
    import cytoscape from 'cytoscape';
    import dagre from 'cytoscape-dagre';
    import klay from 'cytoscape-klay';
    import { onMount } from "svelte";
    import {OpenFiles} from '../../wailsjs/go/main/App'
 
    // cytoscape.use(dagre)
    cytoscape.use(klay)
    // let container: HTMLDivElement;
    let cy: cytoscape.Core

    onMount(() => {
        cy = createCytoscape()
    })

    function createCytoscape(elements: string = '[]'): cytoscape.Core {
        cy?.destroy()
        debugger
        return cytoscape({
            container: document.getElementById('container'),

            elements: JSON.parse(elements),
        
            layout: {
                // name: 'dagre',
                // animate: true,
                // nodeDimensionsIncludeLabels: true,
                // idealEdgeLength: 20,
                // align: 'UL',
                // rankDir: 'TB',
                //
                
                name: 'klay',
                animate: true,
                nodeDimensionsIncludeLabels: true,
                klay: {
                    // Following descriptions taken from http://layout.rtsys.informatik.uni-kiel.de:9444/Providedlayout.html?algorithm=de.cau.cs.kieler.klay.layered
                    addUnnecessaryBendpoints: true, // Adds bend points even if an edge does not change direction.
                    aspectRatio: 2, // The aimed aspect ratio of the drawing, that is the quotient of width by height
                    borderSpacing: 20, // Minimal amount of space to be left to the border
                    compactComponents: true, // Tries to further compact components (disconnected sub-graphs).
                    crossingMinimization: 'LAYER_SWEEP', // Strategy for crossing minimization.
                    /* LAYER_SWEEP The layer sweep algorithm iterates multiple times over the layers, trying to find node orderings that minimize the number of crossings. The algorithm uses randomization to increase the odds of finding a good result. To improve its results, consider increasing the Thoroughness option, which influences the number of iterations done. The Randomization seed also influences results.
                    INTERACTIVE Orders the nodes of each layer by comparing their positions before the layout algorithm was started. The idea is that the relative order of nodes as it was before layout was applied is not changed. This of course requires valid positions for all nodes to have been set on the input graph before calling the layout algorithm. The interactive layer sweep algorithm uses the Interactive Reference Point option to determine which reference point of nodes are used to compare positions. */
                    cycleBreaking: 'GREEDY', // Strategy for cycle breaking. Cycle breaking looks for cycles in the graph and determines which edges to reverse to break the cycles. Reversed edges will end up pointing to the opposite direction of regular edges (that is, reversed edges will point left if edges usually point right).
                    /* GREEDY This algorithm reverses edges greedily. The algorithm tries to avoid edges that have the Priority property set.
                    INTERACTIVE The interactive algorithm tries to reverse edges that already pointed leftwards in the input graph. This requires node and port coordinates to have been set to sensible values.*/
                    direction: 'UNDEFINED', // Overall direction of edges: horizontal (right / left) or vertical (down / up)
                    /* UNDEFINED, RIGHT, LEFT, DOWN, UP */
                    edgeRouting: 'POLYLINE', // Defines how edges are routed (POLYLINE, ORTHOGONAL, SPLINES)
                    edgeSpacingFactor: 0.5, // Factor by which the object spacing is multiplied to arrive at the minimal spacing between edges.
                    feedbackEdges: false, // Whether feedback edges should be highlighted by routing around the nodes.
                    fixedAlignment: 'NONE', // Tells the BK node placer to use a certain alignment instead of taking the optimal result.  This option should usually be left alone.
                    /* NONE Chooses the smallest layout from the four possible candidates.
                    LEFTUP Chooses the left-up candidate from the four possible candidates.
                    RIGHTUP Chooses the right-up candidate from the four possible candidates.
                    LEFTDOWN Chooses the left-down candidate from the four possible candidates.
                    RIGHTDOWN Chooses the right-down candidate from the four possible candidates.
                    BALANCED Creates a balanced layout from the four possible candidates. */
                    inLayerSpacingFactor: 1.0, // Factor by which the usual spacing is multiplied to determine the in-layer spacing between objects.
                    layoutHierarchy: false, // Whether the selected layouter should consider the full hierarchy
                    linearSegmentsDeflectionDampening: 0.3, // Dampens the movement of nodes to keep the diagram from getting too large.
                    mergeEdges: false, // Edges that have no ports are merged so they touch the connected nodes at the same points.
                    mergeHierarchyCrossingEdges: true, // If hierarchical layout is active, hierarchy-crossing edges use as few hierarchical ports as possible.
                    nodeLayering:'LONGEST_PATH', // Strategy for node layering.
                    /* NETWORK_SIMPLEX This algorithm tries to minimize the length of edges. This is the most computationally intensive algorithm. The number of iterations after which it aborts if it hasn't found a result yet can be set with the Maximal Iterations option.
                    LONGEST_PATH A very simple algorithm that distributes nodes along their longest path to a sink node.
                    INTERACTIVE Distributes the nodes into layers by comparing their positions before the layout algorithm was started. The idea is that the relative horizontal order of nodes as it was before layout was applied is not changed. This of course requires valid positions for all nodes to have been set on the input graph before calling the layout algorithm. The interactive node layering algorithm uses the Interactive Reference Point option to determine which reference point of nodes are used to compare positions. */
                    nodePlacement:'BRANDES_KOEPF', // Strategy for Node Placement
                    /* BRANDES_KOEPF Minimizes the number of edge bends at the expense of diagram size: diagrams drawn with this algorithm are usually higher than diagrams drawn with other algorithms.
                    LINEAR_SEGMENTS Computes a balanced placement.
                    INTERACTIVE Tries to keep the preset y coordinates of nodes from the original layout. For dummy nodes, a guess is made to infer their coordinates. Requires the other interactive phase implementations to have run as well.
                    SIMPLE Minimizes the area at the expense of... well, pretty much everything else. */
                    randomizationSeed: 1, // Seed used for pseudo-random number generators to control the layout algorithm; 0 means a new seed is generated
                    routeSelfLoopInside: false, // Whether a self-loop is routed around or inside its node.
                    separateConnectedComponents: true, // Whether each connected component should be processed separately
                    spacing: 20, // Overall setting for the minimal amount of space to be left between objects
                    thoroughness: 100 // How much effort should be spent to produce a nice layout..
                },
            },
            wheelSensitivity: 0.2,
            motionBlur: true,
            // maxZoom: 100,
        
            style: [
                {
                    "selector": "node",
                    "style": {
                        "label": "data(name)",
                        'font-size': '14px',
                        'text-wrap': 'wrap',
                        'border-width': '0px',
                        "text-valign": "top",
                        "text-halign": "left",
                        "width": 0.1,
                        "height": 0.1,
                        "background-color": 'data(bg)',
                        "text-background-color": "blue",
                        'text-margin-x': 60,
                    }
                },
                {
                    "selector": "edge",
                    "style": {
                        'curve-style': 'bezier',
                        'target-arrow-shape': 'circle-triangle',
                        'target-arrow-color': 'black',
                        'line-color': 'black',
                        'width': 2,
                    }
                }
            ],
        });
    }
    async function a() {
        const content = await OpenFiles();
        createCytoscape(content)
    }
</script>

<div>
    <div class="input--button">
        <button class="input--button" on:click={a}>upload</button>
    </div>
    <div id="container"></div>
</div>

<style>
    #container {
            position: fixed;
            left: 0;
            top: 0;
            width: 100%;
            height: 100%;
            z-index: 50;
    }

    .input--button {
        position: fixed;
        z-index: 999;
        cursor: pointer;
    }
</style>