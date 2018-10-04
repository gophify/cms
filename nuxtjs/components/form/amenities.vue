<template>
  <div class="field_wrapper">
    <v-autocomplete class="input_auto_H"
        v-model="data"
        :items="amenities"
        solo
        flat
        chips
        color="blue-grey lighten-2"
        item-text="label"
        item-value="key"
        multiple
        >
        <template
        slot="selection"
        slot-scope="data"
        >
        <v-chip
            :selected="data.selected"
            close
            class="chip--select-multi"
            @input="remove(data.item)"
        >
            <v-avatar>
            <img :src="data.item.icon">
            </v-avatar>
            {{ data.item.label }}
        </v-chip>
        </template>
        <template
        slot="item"
        slot-scope="data"
        >
        <template v-if="typeof data.item !== 'object'">
            <v-list-tile-content v-text="data.item"></v-list-tile-content>
        </template>
        <template v-else>
            <v-list-tile-avatar>
            <img :src="data.item.icon">
            </v-list-tile-avatar>
            <v-list-tile-content>
            <v-list-tile-title v-html="data.item.label"></v-list-tile-title>
            <v-list-tile-sub-title v-html="data.item.group"></v-list-tile-sub-title>
            </v-list-tile-content>
        </template>
        </template>
    </v-autocomplete>
  </div>
</template>

<script>
  export default {
    props: { data: Array, model: String },
    data () {
      let icons = {
        1: this.$store.state.mediaPath + '/assets/images/amenities_icons/ac.png',
        2: this.$store.state.mediaPath + '/assets/images/amenities_icons/bathroom.png',
        3: this.$store.state.mediaPath + '/assets/images/amenities_icons/boat.png',
        4: this.$store.state.mediaPath + '/assets/images/amenities_icons/cleaning.png',
        5: this.$store.state.mediaPath + '/assets/images/amenities_icons/coffee.png',
        6: this.$store.state.mediaPath + '/assets/images/amenities_icons/coworking.png',
        7: this.$store.state.mediaPath + '/assets/images/amenities_icons/kitchen.png',
        8: this.$store.state.mediaPath + '/assets/images/amenities_icons/pool.png',
        9: this.$store.state.mediaPath + '/assets/images/amenities_icons/sheets.png',
        10: this.$store.state.mediaPath + '/assets/images/amenities_icons/toiletry.png',
        11: this.$store.state.mediaPath + '/assets/images/amenities_icons/wifi.png'
      }
      return {
        amenities: [
          { header: 'Group 1' },
          { key: 'ac', label: 'Air conditioning', group: 'Group 1', icon: icons[1] },
          { key: 'bathroom', label: 'Private room with en-suite bathrooms', group: 'Group 1', icon: icons[2] },
          { key: 'boat', label: 'Complementary on-demand boat service', group: 'Group 1', icon: icons[3] },
          { key: 'cleaning', label: 'Cleaning 2x a week with towel change', group: 'Group 1', icon: icons[4] },
          { divider: true },
          { header: 'Group 2' },
          { key: 'coffee', label: 'Unlimited coffee, tea and water', group: 'Group 2', icon: icons[5] },
          { key: 'coworking', label: 'Unlimited access to the coworking space ', group: 'Group 2', icon: icons[6] },
          { key: 'kitchen', label: 'Community Kitchen', group: 'Group 2', icon: icons[7] },
          { key: 'pool', label: 'Private swimming pool', group: 'Group 2', icon: icons[8] },
          { key: 'sheets', label: 'Weekly sheet change', group: 'Group 2', icon: icons[9] },
          { key: 'toiletry', label: 'Toiletry kit', group: 'Group 2', icon: icons[10] },
          { key: 'wifi', label: 'Powerful internet', group: 'Group 2', icon: icons[11] }
        ]
      }
    },
    watch: {
      data: function (val, oldVal) {
        this.emitFormVal()
      }
    },
    methods: {
      remove (item) {
        const index = this.data.indexOf(item.key)
        if (index >= 0) this.data.splice(index, 1)
      },
      emitFormVal () {
        var data = {model: this.model, data: this.data}
        this.$emit('updateFormVal', data)
      }
    }
  }
</script>

<style>
.v-chip .v-chip__content{
  white-space: normal!important;
  height: auto!important;
}
</style>
