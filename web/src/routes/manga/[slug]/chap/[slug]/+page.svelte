<script>
  import ChapterNav from "$lib/components/ChapterNav.svelte";

  // todo: retrieve images from backend
  const images = [1, 2, 3, 4, 5];
</script>

<ChapterNav />

<div class="container">
  {#each images as image}
    <figure class="image">
      <img src="https://bulma.io/assets/images/placeholders/128x128.png" />
    </figure>
  {/each}
</div>
