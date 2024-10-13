<script>
  // should generate index.html
  import Button from './_components/button.svelte'
  import Links from './_components/links.svelte'
  
  let a = '#{title}'
  let b = JSON.stringify({/* obj1 */});
  let c = JSON.stringify([/* arr */]);
  let d = '/*! str1 */';

  function doSomething() {
	  a = 'something'
	  fetch('', {}).then(() => {

	  }).catch(function(){

	  })
  }
</script>

{a}, {b}, {c}, {d}
<button onclick="doSomething()">Something</button>
<Button title="something"></Button>
<Links></Links>
