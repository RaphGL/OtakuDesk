<script lang="ts">
  import Pagination from "$lib/helpers/Pagination.svelte";
  import ContentCard from "$lib/components/ContentCard.svelte";

  type Props = {
    mediaType: "anime" | "manga";
  };
  const { mediaType }: Props = $props();

  type AnimeItem = {
    name: string;
    path: string;
    episode: number;
    curr_episode: number;
    total_episodes: number;
    chapter: number;
    curr_chapter: number;
    total_chapters: number;
  };

  let currPage = $state(0);
  let maxPage = $state(0);
  const ITEMS_PER_PAGE = 15;

  async function loadMediaList(): Promise<AnimeItem[]> {
    let endpoint = "";
    if (mediaType === "anime") {
      endpoint = "http://localhost:8080/animes";
    } else if (mediaType === "manga") {
      endpoint = "http://localhost:8080/mangas";
    } else {
      console.error("invalid media type:", mediaType);
    }
    let resp = await fetch(endpoint);
    let animes: AnimeItem[] = await resp.json();
    maxPage = animes.length / ITEMS_PER_PAGE;
    return animes;
  }

  let animes: AnimeItem[] = $state([]);
  $effect(() => {
    loadMediaList().then((mediaList) => (animes = mediaList));
  });
</script>

<div class="hero">
  <div class="hero-body">
    <h1 class="title has-text-centered">
      {#if mediaType === "anime"}
        Anime
      {:else}
        Manga
      {/if}
      Catalog
    </h1>
  </div>
</div>

<div class="my-6">
  <div class="container fixed-grid has-5-cols container">
    <div class="grid">
      {#if animes}
        {#each animes.slice(currPage * ITEMS_PER_PAGE, currPage * ITEMS_PER_PAGE + ITEMS_PER_PAGE) as anime}
          <div class="cell">
            <ContentCard title={anime.name} href={anime.path} />
          </div>
        {/each}
      {/if}
    </div>
  </div>

  {#if maxPage > 1}
    <Pagination bind:currPage bind:maxPage />
  {/if}
</div>
