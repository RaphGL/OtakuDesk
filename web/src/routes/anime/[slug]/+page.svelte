<script lang="ts">
  import EpisodeList from "$lib/EpisodeList.svelte";
  import ContentInfo from "$lib/ContentInfo.svelte";

  // todo: fetch anime from url path and database
</script>

<div class="container separated-content py-6">
  <!-- video section -->
  <div class="columns">
    <EpisodeList />
    <video width="80%" controls>
      <source src="movie.mp4" type="video/mp4" />
      <source src="movie.ogg" type="video/ogg" />
      Your browser does not support the video tag.
    </video>
  </div>
  <ContentInfo />
</div>

<style>
  video {
    border-radius: 0.4em;
  }

  .separated-content {
    display: flex;
    flex-direction: column;
    gap: 2em;
  }
</style>
