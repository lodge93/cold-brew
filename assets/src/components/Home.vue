<template>
  <v-container fluid>
    <v-slide-y-transition mode="out-in">
      <v-layout row wrap>
        <manual-controls></manual-controls>
      </v-layout>
    </v-slide-y-transition>
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