<script>
  import 'bulma/css/bulma.css';
  import NavBar from '$lib/NavBar.svelte';
  import Footer from '$lib/Footer.svelte';
</script>

<NavBar />
<slot />
<Footer />
