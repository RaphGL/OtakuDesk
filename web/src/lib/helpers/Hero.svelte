<script lang>
  const { 
    children, 
  } = $props();
</script>

<div class="hero is-small">
  <div class="hero-body">
    {@render children()}
  </div>
</div>
