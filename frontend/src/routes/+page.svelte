<script>
    import { username } from '$lib/stores/username';
    import { goto } from '$app/navigation';
</script>

<h1>Hello {$username}!</h1>
<button on:click={() => goto('/login')}>/login</button>
<button on:click={() => goto('/register')}>/register</button>
