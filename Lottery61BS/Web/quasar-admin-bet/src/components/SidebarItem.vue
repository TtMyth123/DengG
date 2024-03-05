<template>
    <div>
        <template
                v-if="!hasChild(item)">
            <essential-link :link="item.path" :caption="item.title" :icon="item.icon"> </essential-link>
        </template>
      <template v-else>
        <q-expansion-item
          :icon="item.icon"
          :label="item.title">
          <q-list class="q-pl-lg">
            <sidebar-item v-for="child in item.children"
                          :key="child.path"
                          :is-nest="true"
                          :item="child"
            />
          </q-list>
        </q-expansion-item>

      </template>
    </div>
</template>

<script>
    import EssentialLink from './EssentialLink';

    export default {
        name: 'SidebarItem',
        components: { EssentialLink },
        props: {
            // route object
            item: {
                type: Object,
                required: true
            },
            isNest: {
                type: Boolean,
                default: false
            },
            basePath: {
                type: String,
                default: ''
            }
        },
        data() {
            // To fix https://github.com/PanJiaChen/vue-admin-template/issues/237
            // TODO: refactor with render function
            this.onlyOneChild = null;
            return {};
        },
        methods: {
            hasChild(item){
              if (item.children==undefined || item.children ==null && item.children.length==0) {
                return false
              }
              return true
            },
        }
    };
</script>
