<script>
  import ContentInfo from "$lib/components/ContentInfo.svelte";
  import ChapterList from "$lib/components/ChapterList.svelte";
  import Pagination from "$lib/helpers/Pagination.svelte";
</script>

<div class="container my-6">
  <ContentInfo />
  <div class="my-6">
    <ChapterList />
    <Pagination />
  </div>
</div>
