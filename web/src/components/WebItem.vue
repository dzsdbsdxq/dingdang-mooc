<template>
  <div class="web-item">
    <h4 class="text-gray">
      <i class="linecons-tag"></i>{{item.name}}
    </h4>
    <div class="row" v-if="item.lists.length > 0">
      <div class="col-sm-3" v-for="(web, idx) in item.lists" :key="idx">
        <div @click="alertVideoDialog(web.path)" class="folder-grid xe-widget xe-conversations box2 label-info">
          <div class="contain">
            <div class="movie-img">
              <img :src="web.logo" alt="streamlab-image" fit="fill">
            </div>
          </div>
          <div class="info">
            <div class="gen-movie-info">
              <p>{{web.title}}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <br />
  </div>
</template>

<script setup>
const props = defineProps({
  item: {
    type:Object,
    default:{}
  },
})
const em = defineEmits(["_alertVideoDialog"])
const alertVideoDialog = (path) => {
  em("_alertVideoDialog",{path:path})
}
</script>

<style scoped lang="less">
.web-item{
  .row{
    .folder-grid{
      display: inline-block;
      width: 100%;
      background-color: #fff;
      border-radius: 8px;
      cursor: pointer;
      .contain{
        position: relative;
        height: 120px;
        margin-bottom: 8px;
        border-radius: 8px 8px 0 0;
        overflow: hidden;
        .movie-img{
          position: relative;
          overflow: hidden;
          &:before{
            content: "";
            opacity: 0;
            z-index: 9;
            position: absolute;
            bottom: 0;
            left: 0;
            width: 100%;
            height: 100%;
            display: inline-block;
            transition: all .5s ease-in-out;
            transition: all .5s ease-in-out;
            -moz-transition: all .5s ease-in-out;
            -ms-transition: all .5s ease-in-out;
            -o-transition: all .5s ease-in-out;
            -webkit-transition: all .5s ease-in-out;
            background: #0B101766;
            background: -webkit-linear-gradient(top,#0B101766,#0B1017CC);
            background: -moz-linear-gradient(top,#0B101766,#0B1017CC);
            background: linear-gradient(to bottom,#0B101766,#0B1017CC);
          }
          img{
            max-width: 100%;
            height: auto;
          }
        }
      }
      .info{
        display: inline-block;
        width: 100%;
        .gen-movie-info{
          p{
            font-size: 16px;
            overflow: hidden;
            text-overflow: ellipsis;
            display: -webkit-box;
            -webkit-line-clamp: 2;
            -webkit-box-orient: vertical;
            color: #545c63;
            line-height: 20px;
            height: 40px;
            margin-bottom: 8px;
            padding: 0 8px;
          }

        }
      }
    }

  }
}
</style>
