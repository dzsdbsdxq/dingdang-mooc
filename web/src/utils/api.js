import {fly} from './request'


export const API = {
    getIndexData: (paramObj) => fly.post('/index', paramObj),
    getListsData:(paramObj) => fly.post('/lists', paramObj),
    getPlayUrl:(paramObj) => fly.post('/player', paramObj),
    getFileInfo:(paramObj) => fly.post('/api/public/getFile', paramObj),
    getDocUrl:(paramObj) => fly.post('/api/public/getDocUrl', paramObj),
    getDownLoadUrl:(paramObj) => fly.post('/api/public/getDownloadUrl', paramObj),

    getCommon: (paramObj) => fly.post('/copyright', paramObj),
    getFileDetail:(paramObj) => fly.post('/api/public/detail', paramObj),

    getAppInfo: (paramObj) => fly.post('/api/public/getAppInfo', paramObj),
    getFiles:(paramObj) => fly.post('/api/public/getFiles', paramObj),
    getLoginQr:(paramObj)=>fly.post("/api/public/getLoginQRCode",paramObj)
}

export const ALL_POST = {
    allRequest:(paramsArrs) => fly.all(paramsArrs)
}
