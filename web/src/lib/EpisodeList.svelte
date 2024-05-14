<script>
  let { currEpisode = $bindable(1) } = $props();

  // todo: change with fetching episodes from database
  const noOfEpisodes = [];
  for (let i = 1; i < 10; i++) {
    noOfEpisodes.push(i);
  }
</script>

<div class="panel m-2 p-6">
  {#each noOfEpisodes as episode}
    <p class="panel-block">
      <button
        onclick={() => (currEpisode = episode)}
        class="episode-link"
        class:has-text-success={currEpisode === episode}
        >Episode {episode}</button
      >
    </p>
  {/each}
</div>

<style>
  .episode-link:hover {
    color: var(--bulma-primary);
  }
</style>
