<template>
  <div id="app">
    <div class="content">
      <div class="server_tab">
        <el-card class="box-card" v-for="server in serverList" :key="server.IP">
          <div slot="header" class="clearfix">
            <span><img :src="server.Flag" width="20" alt="Flag" /></span>
            <span>{{ server.ISP }}</span>
            <span style="float: right"> </span>
            <el-popover
              style="float: right"
              placement="top-start"
              title="详情数据"
              width="25vw"
              max-width="500px"
              trigger="hover"
              content=""
            >
              <div class="popover-content" v-html="serverDetailInfo"></div>
              <i
                class="el-icon-s-promotion"
                slot="reference"
                style="float: right; color: #243f59; vertical-align: center"
              ></i>
            </el-popover>
          </div>
          <el-row class="row">
            <el-col :span="4">cpu</el-col>
            <el-col :span="20"
              ><el-progress
                :text-inside="true"
                :stroke-width="26"
                :percentage="Math.floor(server.CPUUsed * 100) / 100"
                :status="setItemStatus(server.CPUUsed)"
              ></el-progress>
            </el-col>
          </el-row>
          <el-row class="row">
            <el-col :span="4">内存</el-col>
            <el-col :span="20"
              ><el-progress
                :text-inside="true"
                :stroke-width="26"
                :percentage="Math.floor(server.MemUsed * 100) / 100"
                :status="setItemStatus(server.MemUsed)"
              ></el-progress>
            </el-col>
          </el-row>
          <el-row class="row">
            <el-col :span="4">网速</el-col>
            <el-col :span="20">
              <img src="./assets/download.svg" width="17px" height="17px" />
              <span style="line-height: 26px"> {{ server.RecvTraffic }}</span>
              <img src="./assets/upload.svg" width="17px" height="17px" />
              {{ server.SendTraffic }}
            </el-col>
          </el-row>
          <el-row class="row">
            <el-col :span="4">系统</el-col>
            <el-col :span="20" class="col_txt">
              {{ server.Os }}/{{ server.Platform }}
            </el-col>
          </el-row>
          <el-row class="row">
            <el-col :span="4">cpu架构</el-col>
            <el-col :span="20" class="col_txt">
              {{ server.ModelName }}
            </el-col>
          </el-row>
          <el-row class="row">
            <el-col :span="4">在线</el-col>
            <el-col :span="20" class="col_txt">
              {{ server.Uptime }}
            </el-col>
          </el-row>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "App",
  el: "#app",
  data: function () {
    return {
      visible: false,
      serverList: [],
      serverDetailInfo: "",
    };
  },
  created: function () {
    var ws = new WebSocket("ws://" + location.host+ "/ping");
    //连接打开时触发
    ws.onopen = function (evt) {
      console.log("Connection open ...", evt.data);
      ws.send("Hello WebSockets!");
    };
    //接收到消息时触发
    ws.onmessage = this.onMsg;
    // ws.onmessage = function (evt) {
    //    console.log("Received Message: " + evt.data);
    // };
    //连接关闭时触发
    ws.onclose = function (evt) {
      console.log("Connection closed.", evt.data);
    };
  },
  methods: {
    onMsg: function (evt) {
      this.serverList = JSON.parse(evt.data);
      console.log("Received Message: " + evt.data);
    },
    setItemStatus: function (num) {
      if (num > 80) {
        return "exception";
      } else if (num > 50) {
        return "warning";
      } else {
        return "success";
      }
    },
  },
};
</script>

<style>
.col_txt {
  height: 26px;
  line-height: 26px;
  overflow: hidden;
}
body {
  content: " " !important;
  background: fixed !important;
  z-index: -1 !important;
  top: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  left: 0 !important;
  background-position: top !important;
  background-repeat: no-repeat !important;
  background-size: cover !important;
  background-image: url(./assets/bg2.webp) !important;
}
.row {
  padding-bottom: 0.2rem;
}
.content {
  width: 90vw;
  margin: 0 auto;
  background-color: #fbfbfb26;
  border-radius: 1rem !important;
  padding: 1vw;
  border: 10px;
  min-width: 305px;
}
.server_tab {
  display: inline-block;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both;
}

.box-card {
  color: #ffffff !important;
  display: inline-block;
  width: 22vw;
  min-width: 300px;
  border-radius: 0.5rem;
  margin: 3px;
  background-color: transparent !important;
}
</style>