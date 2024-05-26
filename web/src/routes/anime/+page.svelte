<script>
  import Pagination from "$lib/helpers/Pagination.svelte";
  import ContentCard from '$lib/components/ContentCard.svelte';
</script>

<div class="hero">
  <div class="hero-body">
    <h1 class="title has-text-centered">Anime Catalog</h1>
  </div>
</div>


<div class="my-6">
<div class="container fixed-grid has-5-cols container">
<div class="grid">
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
  <div class="cell"><ContentCard /></div>
</div>
</div>

<Pagination />
</div>
