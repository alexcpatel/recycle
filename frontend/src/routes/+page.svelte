<script>
    import { MapLibre } from "svelte-maplibre";
</script>

<MapLibre
    center={[-74.006, 40.7128]}
    zoom={10}
    class="map"
    standardControls
    style="https://basemaps.cartocdn.com/gl/positron-gl-style/style.json"
/>

<style>
    :global(.map) {
        height: 100vh;
    }
</style>
