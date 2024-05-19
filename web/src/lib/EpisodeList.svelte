<script>
  let { currEpisode = $bindable(1), total = 10 } = $props();

  // todo: change with fetching episodes from database
  const noOfEpisodes = [];
  for (let i = 1; i < total; i++) {
    noOfEpisodes.push(i);
  }
</script>

<div class="card m-2 p-4">
  <div class="card-header">
    <!-- selector to choose an episode range -->
    <div class="select is-small">
      <!-- todo: implement episode range selector -->
      <select>
        <option value="1-10">1-10</option>
        <option value="10-20">10-20</option>
        <option value="20-30">20-30</option>
      </select>
    </div>

    <!-- input box to go to specific episode -->
    <div class="field has-addons">
      <div class="control">
        <input
          class="input is-small"
          size="8"
          type="text"
          placeholder="Ep number"
        />
      </div>
      <div class="control">
        <button class="button is-small is-success">goto</button>
      </div>
    </div>
  </div>

  <div class="card-content">
    {#each noOfEpisodes as episode}
      <button
        onclick={() => (currEpisode = episode)}
        class="button is-fullwidth episode-link"
        class:has-text-success={currEpisode === episode}
        >Episode {episode}</button
      >
    {/each}
  </div>
</div>

<style>
  .episode-link {
    border: none;
  }

  .episode-link:hover {
    color: var(--bulma-primar);
  }

  div.card-header {
    display: flex;
    gap: 0.5em;
    justify-content: center;
  }
</style>
