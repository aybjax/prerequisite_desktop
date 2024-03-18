<script lang="ts">
    // @ts-ignore
    import cytoscape from 'cytoscape';
    import dagre from 'cytoscape-dagre';
    import { onMount } from "svelte";
    import {OpenFiles} from '../../wailsjs/go/main/App'
    import Tooltip, { Wrapper }  from '@smui/tooltip'
    import IconButton, { Icon } from '@smui/icon-button';
    import { mdiDownload } from '@mdi/js';
    import Snackbar, { Label as SnackLabel, Actions } from '@smui/snackbar';
    import {UUID} from 'uuidjs'
    

    let nodeFirst: cytoscape.NodeDataDefinition |undefined;
 
    cytoscape.use(dagre);
    let container: HTMLDivElement;
    let errorSnackbar: Snackbar;
    let errorMsg = ''
    let cy: cytoscape.Core

    onMount(() => {
        cy = createCytoscape()
    })

    function createCytoscape(elements: string = '[]'): cytoscape.Core {
        cy?.destroy()
        cy = cytoscape({
            container: container,

            elements: JSON.parse(elements),

            layout: {
                //@ts-ignore
                name: 'dagre',
                // @ts-ignore
                animate: false,
                nodeDimensionsIncludeLabels: true,
                //@ts-ignore
                idealEdgeLength: 100,
                // Extra spacing between components in non-compound graphs
                componentSpacing: 40,

                // Node repulsion (non overlapping) multiplier
                //@ts-ignore
                nodeRepulsion: function( node ){ return 2048; },

                // Node repulsion (overlapping) multiplier
                nodeOverlap: 4,

                // Divisor to compute edge forces
                //@ts-ignore
                edgeElasticity: function( edge ){ return 100; },
            },
            wheelSensitivity: 0.2,
            motionBlur: true,
            // maxZoom: 100,

            style: [
                {
                    "selector": "node",
                    "style": {
                        "label": "data(name)",
                        'font-size': '18px',
                        'text-wrap': 'wrap',
                        'border-width': '0px',
                        "text-valign": "top",
                        "text-halign": "center",
                        "width": 100,
                        "height": 100,
                        "text-background-color": "rgb(255,255,255)",
                        "text-background-opacity": 1,
                        "text-opacity": 1,
                        "opacity": 0.7,
                        'text-margin-x': 0,
                        'text-margin-y': -15,
                        "z-index": 100,
                        "text-max-width": "200",
                    }
                },
                {
                    "selector": "edge",
                    "style": {
                        'curve-style': 'bezier',
                        'target-arrow-shape': 'triangle',
                        'width': 5
                    }
                }
            ]
        });
        cy?.edges().on('dbltap', edgeDbTap)
        cy.on('tap', bgTap);
        
        cy?.nodes().on('dbltap', nodeDbTap)

        return cy;
    }
    async function a() {
        debugger
        let content = '[]';
        try {
            content = await OpenFiles();
        } catch(e) {
            errorMsg = e;
            errorSnackbar.open();
        }
        createCytoscape(content)
    }


    function bgTap(event: cytoscape.EventObject){
    if( event.target !== cy ) return;
        nodeFirst = undefined;
    }
    function edgeDbTap(e: cytoscape.EventObject){
        nodeFirst = undefined
        e.target.remove()
    }
    function nodeDbTap(e: cytoscape.EventObject) {
        if (!nodeFirst) {
            nodeFirst = e.target
            return
        }
        const existingNode = nodeFirst
        nodeFirst = undefined
        if(existingNode === e.target) {
            setTimeout(() => {
            cy.$(`#${existingNode.data('id')}`).unselect()
        }, 200)
        return
    }

    //@ts-ignore
    if(cy.edges().some(edge => (edge.data('source') === existingNode?.data('id') && edge.data('target') === e.target.data('id'))
    //@ts-ignore
            || (edge.data('target') === existingNode?.data('id') && edge.data('source') === e.target.data('id')))) {
                return
            }

    const id = UUID.generate()
    cy.add({
        data: {
        id,
        source: existingNode?.data('id'),
        target: e.target.data('id'),
        }
    })

    cy.$(`#${id}`).on('dbltap', edgeDbTap)
    }


</script>

<div class="aybjax">
    <div class="input--button">
        <!-- <button class="input--button" on:click={a}>upload</button> -->
        <Wrapper>
            <IconButton class="material-icons"  on:click={a}>
                <Icon tag="svg" viewBox="0 0 24 24">
                    <path fill="currentColor" d={mdiDownload} />
                </Icon>
            </IconButton>
            <Tooltip unbounded>Загрузить данные для отображения</Tooltip>
        </Wrapper>
         
        <Snackbar bind:this={errorSnackbar} class="demo-error">
            <SnackLabel>
                {errorMsg}
            </SnackLabel>
            <Actions>
                <IconButton class="material-icons" title="Dismiss">close</IconButton>
            </Actions>
        </Snackbar>
    </div>
    <div id="container" bind:this={container}></div>
</div>

<style lang="scss">
    // See https://github.com/material-components/material-components-web/tree/v14.0.0/packages/mdc-theme
    @use '../../node_modules/@material/theme/color-palette';
    @use '../../node_modules/@material/theme/theme-color';
    // Make sure SMUI's import happens first, since it specifies variables.
    @use '../../node_modules/@smui/snackbar/style' as smui-snackabar;
    // See https://github.com/material-components/material-components-web/tree/v14.0.0/packages/mdc-snackbar
    @use '../../node_modules/@material/snackbar/mixins' as snackbar;
    
    .aybjax {
        #container {
                position: fixed;
                left: 0;
                top: 0;
                width: 100%;
                height: 100%;
                z-index: 50;
        }
    }
    .mdc-snackbar.demo-error {
        @include snackbar.fill-color(color-palette.$red-500);
        @include snackbar.label-ink-color(
            theme-color.accessible-ink-color(color-palette.$red-500)
        );
    }
    .input--button {
        position: fixed;
        z-index: 999;
        cursor: pointer;
    }
</style>