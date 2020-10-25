<template>
  <p>{{ greet }}</p>
</template>
<script>
export default {
  async asyncData({ $axios }) {
    try {
      const greet = await $axios.$get('http://rp:1234/api/greeting')
      return { greet }
    } catch (e) {
      console.error(e)
    }
  }
}
</script>
