<script lang="ts">
  let { currEpisode = $bindable(1) } = $props();

  // todo: change with fetching episodes from database
  type Episode = {
    episode: number;
    released: string;
  };

  const episodes: Episode[] = [];
  for (let i = 1; i < 10; i++) {
    episodes.push({ episode: i, released: "Sep 14, 2020" });
  }
</script>

<table class="table is-hoverable">
  <thead>
    <tr>
      <th class="has-text-left">Chapter</th>
      <th class="has-text-right">Released</th>
    </tr>
  </thead>
  <tbody>
    {#each episodes as { episode, released }}
      <tr>
        <td class="has-text-left">
          <a
            onclick={() => (currEpisode = episode)}
            class="chapter-link"
            class:has-text-success={currEpisode === episode}
            >Episode {episode}</a
          >
        </td>
        <td class="has-text-right">{released}</td>
      </tr>
    {/each}
  </tbody>
</table>

<style>
  .chapter-link:hover {
    color: var(--bulma-primary);
  }

  table {
    width: 90%;
    margin: auto;
  }
</style>
