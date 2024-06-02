<!-- anime/manga card with name and episode number -->
<script>
  const {
    title = "Anime Name",
    subtitle = "",
    href = "#", // link for the anime/manga page
    imgsrc = "#", // thumbnail link
  } = $props();
</script>

<div class="card block p-2">
  <a {href} class="card-image">
    <figure class="image is-4by3" class:is-skeleton={imgsrc === "#"}>
      <img alt="{title} - {subtitle}" src={imgsrc} />
    </figure>
  </a>

  <div class="card-content">
    <div class="has-text-centered is-flex is-flex-direction-column">
      <a {href} class="title is-5">{title}</a>
      {#if subtitle}
        <a {href} class="subtitle is-6">{subtitle}</a>
      {/if}
    </div>
  </div>
</div>

<style>
  .card:not(:hover) {
    box-shadow: none;
  }
</style>
