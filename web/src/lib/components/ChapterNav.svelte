<script>
  import Dropdown from "$lib/helpers/Dropdown.svelte";
  import PrevIcon from "$lib/icons/PrevIcon.svelte";
  import NextIcon from "$lib/icons/NextIcon.svelte";

  const chapters = ["chapter 1", "chapter 2", "chapter 3"];
  let currChap = $state("");
</script>

<div class="box level">
  <div class="level-left">
    <div class="level-item">
      <Dropdown items={chapters} bind:value={currChap} />
    </div>
  </div>

  <div class="level-item">
    <a href={"#"} class="has-text-weight-bold has-text-success">Anime name</a>
  </div>

  <div class="level-right">
    <div class="level-item">
      <div class="pagination" role="navigation" aria-label="pagination">
        <button class="pagination-previous">
          <span class="icon">
            <PrevIcon />
          </span>
        </button>
        <button class="pagination-next">
          <span class="icon">
            <NextIcon />
          </span>
        </button>
      </div>
    </div>
  </div>
</div>
