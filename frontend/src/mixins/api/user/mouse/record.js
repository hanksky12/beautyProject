import {WebSocketMixin} from "@/mixins/utils/websocket/websocket";
import {GetTokenApiMixin} from "@/mixins/api/user/authentication/token";

const ACTIONS = {
    MOVE: "移動",
    CLICK: "點擊",
    SCROLL: "滾動"
};


export const RecordApiMixin = {
    mixins: [
        WebSocketMixin,
        GetTokenApiMixin
    ],
    data() {
        return {
            trackingEvents: {},
        };
    },
    methods: {
        async startRecordApi(selections) {
            console.log("startRecordApi");
            //move,click,scroll
            console.log(selections);
            const action = selections[0].action
            let token = await this.getToken();
            if (!token) {
                console.log('取token失敗')
                return;
            }
            // 開始 WebSocket 連接
            this.startWsConnect(
                action,
                {
                    path: "/user/mouse/tracking",
                    method: "get",
                    body: {
                        action: action,
                        token: token,
                    },
                },
                "已經在監控滑鼠"+action,
                () => {
                    this.startTrackingEvents(action);
                    return  "開始監控滑鼠"+action;
                },
                (data) => {
                    console.log("Message received from WebSocket:", data);
                    return "收到消息";
                },
                (error) => {
                    console.error("WebSocket error:", error);
                    return "連線失敗";
                },
                () => {
                    this.stopTrackingEvents(action);
                    return "停止監控滑鼠"+action;
                },
            );
        },
        async stopRecordApi(selections, killAll=false) {
            console.log("stopRecordApi");
            const action = selections? selections[0].action:""
            this.stopTrackingEvents(action, killAll); // 停止業務邏輯
            this.stopWsConnect( action, "即將結束監控.."+action, action+"沒有在監控",killAll);
        },
        async getToken() {
            const responseData = await this.getTokenApi()
            return responseData.code === 200  ? responseData.data.message: '';
        },
        // 業務邏輯：開始監控硬體使用率
        startTrackingEvents(action) {
            console.log(`Start tracking ${action} event`);
            if (this.trackingEvents && this.trackingEvents[action]) {
                console.warn(`${action} tracking already started.`);
                return;
            }
            const sendMouseData = (event) => {
                const data = {
                    x: event.clientX,
                    y: event.clientY,
                };
                this.sendTrackingData(action, data);
            };
            const sendScrollData = () => {
                const data = {
                    x: window.scrollX,
                    y: window.scrollY,
                };
                this.sendTrackingData(action, data);
            };

            if (action === ACTIONS.MOVE) {
                document.addEventListener("mousemove", sendMouseData);
                this.trackingEvents[action] = sendMouseData;
            } else if (action === ACTIONS.CLICK) {
                document.addEventListener("click", sendMouseData);
                this.trackingEvents[action] = sendMouseData;
            } else if (action === ACTIONS.SCROLL) {
                window.addEventListener("scroll", sendScrollData);
                this.trackingEvents[action] = sendScrollData;
            } else {
                console.warn("Invalid action for tracking.");
            }
        },

        // 業務邏輯：停止監控硬體使用率
        stopTrackingEvents(action, killAll=false) {
            if (killAll) {
                this.killAll();
                return
            }
            if (!this.trackingEvents.active) {
                console.log("No tracking events to stop.");
                return;
            }
            // if (this.trackingEvents === null) {
            //     return
            // }
            if (action === ACTIONS.MOVE || action === ACTIONS.CLICK) {
                document.removeEventListener(action === ACTIONS.MOVE ? "mousemove" : "click", this.trackingEvents[action]);
                delete this.trackingEvents[action];
                console.log(`Stopped tracking ${action} event.`);
            } else if (action === ACTIONS.SCROLL) {
                window.removeEventListener("scroll", this.trackingEvents.scroll);
                delete this.trackingEvents.scroll;
                console.log("Stopped tracking scroll event.");
            } else {
                console.log("Invalid action provided for stopping tracking.");
            }
        },
        sendTrackingData(action, data) {
            const webSocket = this.webSockets[action];
            if (webSocket && webSocket.readyState === webSocket.OPEN) {
                const sendMsg = JSON.stringify(data);
                console.log("Sending tracking data:", sendMsg);
                webSocket.send(sendMsg);
            }
        },
        killAll() {
            for (const action in this.trackingEvents) {
                if (this.trackingEvents[action] && action !== "active") {
                    this.stopTrackingEvents(action);
                }
            }
            this.trackingEvents = {};
            console.log("Stopped all tracking events.");
        },
    },
    beforeDestroy() {
        this.stopRecordApi("noUse",true); // 確保在換頁時 時清理
    },
};