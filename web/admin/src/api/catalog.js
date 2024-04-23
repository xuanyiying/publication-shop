import service from './index'

export const listPublication = (pageNum, pageSize) => {
    return service.get("/admin/v1/catalog/Publications", {
        pageNum, pageSize
    })
};

export const getPublicationDetail = (id) => {
    return service.get("/admin/v1/catalog/Publications" + id)
};
