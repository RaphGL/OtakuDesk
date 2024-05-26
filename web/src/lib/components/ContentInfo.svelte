<script lang="ts">
  // todo: use database/API to retrieve anime info and maybe get a prop for anime ID/name whatever
  const animeName = "Anime Name";
  const animeNameJP = "Anime no Namae";

  type ContentInfo = {
    type: string;
    country: string;
    status: string;
    genres: string[];
    mal: number;
    duration: number;
    episodes: number;
    studios: string[];
    producers: string[];
  };

  const contentInfo: ContentInfo = {
    type: "TV",
    country: "Japan",
    status: "Completed",
    genres: ["shounen", "action", "drama"],
    mal: 8.2,
    duration: 24,
    episodes: 12,
    studios: ["mappa", "kyoani"],
    producers: ["mappa"],
  };
</script>

<!-- section with information about the anime/manga -->
<div class="columns container">
  <div class="column is-one-third is-skeleton">
    <figure class="image is-4by5">
      <img src="anime cover" alt="#" />
    </figure>
  </div>
  <div class="column">
    <p class="title">{animeName}</p>
    <p class="subtitle">{animeNameJP}</p>
    <p>
      As players of Monster Hunter and Dungeons & Dragons know, the slime is not
      exactly the king of the fantasy monsters. So when a 37-year-old Tokyo
      salaryman dies and wakes up in a world of dragons and magic, he's a little
      disappointed to find he's become a blind, boneless slime monster.
    </p>

    <ul class="my-4">
      {#each Object.entries(contentInfo) as [infoName, infoValue]}
        <li>
          <span class="has-text-light">{infoName}:</span>
          <span class="has-text-success">
            {#if infoValue instanceof Array}
              {#each infoValue as value, i}
                <a>{value}{i < infoValue.length - 1 ? "," : ""} </a>
              {/each}
            {:else}
              {infoValue}
            {/if}
          </span>
        </li>
      {/each}
    </ul>
  </div>
</div>

<style>
  a {
    color: var(--bulma-success);
  }

  a:hover {
    color: var(--bulma-link);
  }
</style>
