<template>
  <v-container fluid>
      <v-layout row wrap>
        <manual-controls></manual-controls>
      </v-layout>
  </v-container>
</template>

<script>
import ManualControls from './ManualControls'

export default {
  name: 'Home',
  components: {
    ManualControls
  }
}
</script>