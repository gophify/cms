<template>
    <v-flex class="border_section">
        <span class="icon_section"><v-icon dark small>check</v-icon></span>
        <span class="label_section">{{ label }}</span>
    </v-flex>
</template>

<script>
	export default {
	  props: { label: String }
	}
</script>