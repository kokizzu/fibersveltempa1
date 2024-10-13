<script>
  // should generate index.html
  import Button from './_components/button.svelte'
  import Links from './_components/links.svelte'
  
  let a = '#{title}'
  let b = {/* obj1 */};
  let users = [/* arr */];
  let errMsg = '/*! str1 */';

  function doSomething() {
	  a = 'something'
	  fetch('/get-users', {
		  method: 'POST',
	  }).then(async (res) => {
			users = await res.json()
	  }).catch(function(err){
			errMsg = err
	  })
  }
</script>

{a}, {JSON.stringify(b)}, {errMsg}

<ul>
	{#each users as user}
		<li>{user.Name} {user.Age}</li>
		{/each}
</ul>

<button on:click={() => {doSomething()}}>Something</button>
<Links></Links>
