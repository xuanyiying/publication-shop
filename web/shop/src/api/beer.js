import service from './index'

export const listBook = (pageNum, pageSize) => {
    return service.get("/v1/catalog/Books", {
        pageNum, pageSize
    })
};

export const getBookDetail = (id) => {
    return service.get("/v1/catalog/Books/"+id)
};
