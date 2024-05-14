<script>
  import EpisodeList from '$lib/EpisodeList.svelte';

  // todo: fetch anime from url path and database
  const animeName = "Anime Name";
</script>

<section class="section">
  <p class="title has-text-centered">
  {animeName}
  </p>
</section>

<div class="container">
<div class="columns">
  <EpisodeList />
  <video width="80%" controls>
    <source src="movie.mp4" type="video/mp4" />
    <source src="movie.ogg" type="video/ogg" />
    Your browser does not support the video tag.
  </video>
</div>


<div class="columns">
  <div class="column is-one-third is-skeleton">
    <figure class="image is-4by5">
      <img src="anime cover" alt="#"/>
    </figure>
  </div>
  <p class="columns is-vcentered m-4">
    As players of Monster Hunter and Dungeons & Dragons know, the slime is not
    exactly the king of the fantasy monsters. So when a 37-year-old Tokyo
    salaryman dies and wakes up in a world of dragons and magic, he's a little
    disappointed to find he's become a blind, boneless slime monster.
  </p>
</div>
</div>

<style>
  video {
    border-radius: 0.4em;
  }
</style>

