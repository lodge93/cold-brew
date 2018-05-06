<template>
  <v-container fluid>
      <v-layout row wrap>
        <dripper-settings></dripper-settings>
      </v-layout>
  </v-container>
</template>

<script>
import DripperSettings from './DripperSettings'

export default {
  name: 'Home',
  components: {
    DripperSettings
  }
}
</script>