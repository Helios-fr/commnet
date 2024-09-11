<script>
    import { Button } from '$lib/components/ui/button/index.js';
    import { DefaultUsername } from '$lib/wailsjs/go/main/App.js';

    import { username, isLoggedIn } from '$lib/stores/username';
    import { goto } from '$app/navigation';

    async function logout() {
        console.log('logout')
        await username.logout()
        $isLoggedIn = false
        goto('/')
    }
</script>

<h1 class="text-2xl font-bold my-2">Hello {$username}!</h1>

{#if $isLoggedIn}
    <Button on:click={() => logout()}>Logout</Button>
{:else}
    <Button on:click={() => goto('/login')}>Login</Button>
    <Button on:click={() => goto('/register')}>Register</Button>
{/if}
