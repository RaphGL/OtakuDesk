<script lang="ts">
  import Dropdown from "$lib/helpers/Dropdown.svelte";

  type Props = {
    currEpisode: number;
    total: number;
  };
  let { currEpisode = $bindable(1), total = 10 }: Props = $props();

  // todo: change with fetching episodes from database
  type Episode = {
    ep: number;
    title: string;
  };

  const episodes: Episode[] = [];
  for (let i = 1; i < total; i++) {
    episodes.push({ ep: i, title: `Episode ${i}` });
  }

  let selectedRange = $state("");
</script>

<div class="card m-2 p-4">
  <div class="card-header">
    <!-- selector to choose an episode range -->
    <div class="select is-small">
      <!-- todo: implement episode range selector -->
      <Dropdown
        bind:value={selectedRange}
        size="is-small"
        items={["1-10", "10-20", "20-30"]}
      />
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
    {#each episodes as { ep, title } (ep)}
      <button
        onclick={() => (currEpisode = ep)}
        class="button is-fullwidth episode-link"
        class:has-text-success={currEpisode === ep}>{title}</button
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
