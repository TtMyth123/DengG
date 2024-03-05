<template>
  <q-layout view="lHh lpR fFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          @click="toggleLeftDrawer"
          icon="menu"
          aria-label="Menu"
        />
        <q-toolbar-title>
          {{ $store.state.sys.title }}
        </q-toolbar-title>
        <q-space/>
        <div class="q-gutter-sm row items-center no-wrap">
                    <q-btn round dense flat color="white" :icon="$q.fullscreen.isActive ? 'fullscreen_exit' : 'fullscreen'"
                           @click="$q.fullscreen.toggle()"
                           v-if="$q.screen.gt.sm">
                    </q-btn>

          <q-btn round flat >
            <q-avatar size="26px">
              <img src="https://cdn.quasar.dev/img/boy-avatar.png">
            </q-avatar>
            <q-menu>
              <q-list style="min-width: 100px">
                <q-item clickable v-close-popup to="/UserProfile">
                  <q-item-section>我的信息</q-item-section>
                </q-item>
              </q-list>
              <q-list style="min-width: 100px">
                <q-item clickable v-close-popup @click="onClickLogout">
                  <q-item-section>退出</q-item-section>
                </q-item>
              </q-list>

            </q-menu>
          </q-btn>
        </div>
      </q-toolbar>
    </q-header>

    <!--      show-if-above-->
    <q-drawer
      v-model="leftDrawerOpen"
      bordered
    >
      <q-list>
        <sidebar-item
          v-for="child in getSysMenuList"
          :key="child.path"
          :is-nest="true"
          :item="child"
        >
        </sidebar-item>

      </q-list>
    </q-drawer>

    <q-page-container class="bg-grey-2">
      <router-view/>
    </q-page-container>
  </q-layout>
</template>

<script>
import EssentialLink from 'components/EssentialLink.vue'
import SidebarItem from 'components/SidebarItem.vue'
import Messages from "./Messages.vue";
import { mapState } from 'vuex';

import {defineComponent, ref} from 'vue'
import {useQuasar} from "quasar";
import {Logout } from "src/api/RegLogin";
export default defineComponent({
  name: 'MainLayout',
  components: {
    EssentialLink,
    SidebarItem,
    Messages
  },
  computed: {
    ...mapState({
      getSysMenuList: (state => state.user.sysMenuList),
      }
    )
  },
  data(){
    return{
      RefreshData:{
        refreshDataTimer:null,
        MaxCurResultSecond:2*10,
        CountdownSecond:2*1 ,
        CurResultSecond:0,
      },
    }
  },
  methods:{
    onClickLogout(){
      Logout().then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.$router.push("/Login")
        } else {
          // this.$q.notify({message:msg})
          this.$router.push("/Login")
        }
      }).catch(error => {
        this.$q.notify({message:error})
        this.$router.push("/Login")
      })
    },
    autoRefreshData() {
      this.RefreshData.CurResultSecond++
      if (this.RefreshData.CurResultSecond > 1000000) {
        this.RefreshData.CurResultSecond = 0
      }

      if (this.RefreshData.CurResultSecond % this.RefreshData.CountdownSecond==0) {

      }
    },
    reloadMenu(){
      this.$store.dispatch('user/loadMenu')
    },
  },
  created() {
    this.reloadMenu()
  },
  mounted() {
    if (this.RefreshData.refreshDataTimer) {
      clearInterval(this.RefreshData.refreshDataTimer)
      this.RefreshData.refreshDataTimer = null;
    }
    this.RefreshData.refreshDataTimer = setInterval(()=>{this.autoRefreshData()},500);
  },
  unmounted() {
    clearInterval(this.RefreshData.refreshDataTimer);
    this.RefreshData.refreshDataTimer = null
  },
  setup() {
    const leftDrawerOpen = ref(false)
    // const $q = useQuasar()
    return {
      // $q,
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value
      }
    }
  }
})
</script>

<style>

/* FONT AWESOME GENERIC BEAT */
.fa-beat {
  animation: fa-beat 5s ease infinite;
}

@keyframes fa-beat {
  0% {
    transform: scale(1);
  }
  5% {
    transform: scale(1.25);
  }
  20% {
    transform: scale(1);
  }
  30% {
    transform: scale(1);
  }
  35% {
    transform: scale(1.25);
  }
  50% {
    transform: scale(1);
  }
  55% {
    transform: scale(1.25);
  }
  70% {
    transform: scale(1);
  }
}

</style>
