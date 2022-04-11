import {NzTableQueryParams} from "ng-zorro-antd/table";


export function parseTableQuery(query: NzTableQueryParams, body: any): void {

  //过滤器
  body.filter = {}
  query.filter.forEach(f => {
      body.filter[f.key] = f.value;
  })

  //分布
  body.skip = (query.pageIndex - 1) * query.pageSize;
  body.limit = query.pageSize;

  //排序
  const sorts = query.sort.filter(s => s.value);
  if (sorts.length) {
    body.sort = {};
    sorts.forEach(s => {
      body.sort[s.key] = s.value === 'ascend' ? 1 : -1;
    });
  } else {
    delete body.sort;
  }
}
