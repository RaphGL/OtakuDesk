<script>
  import ContentCard from "$lib/ContentCard.svelte";
  import Hero from "$lib/helpers/Hero.svelte";
</script>

{#snippet contents()}
<div class="columns is-col-min-9 box is-shadowless mx-5 x-scrollable">
  <div class="column"><ContentCard /></div>
  <div class="column"><ContentCard /></div>
  <div class="column"><ContentCard /></div>
  <div class="column"><ContentCard /></div>
  <div class="column"><ContentCard /></div>
  <div class="column"><ContentCard /></div>
  <div class="column"><ContentCard /></div>
  <div class="column"><ContentCard /></div>
</div>
{/snippet}

{#snippet section_title(title, href)}
<Hero>
  <p class="title">{title} <a {href} class="is-size-6 has-text-weight-light has-text-success" >View More</a></p>
</Hero>
{/snippet}

{@render section_title('Recently Watched', '#')}
{@render contents()}

{@render section_title('Recently Read', '#')}
{@render contents()}

<style>
  .x-scrollable {
    overflow-x: scroll;
  }
</style>
