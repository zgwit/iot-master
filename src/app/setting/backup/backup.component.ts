import { Component } from '@angular/core';

@Component({
  selector: 'app-backup',
  templateUrl: './backup.component.html',
  styleUrls: ['./backup.component.scss']
})
export class BackupComponent {
  href!:string
  list:any[]=['id','名称','描述','操作']
  listOfData: any[] = [
    {
      id: '1',
      name: '页面备份',
      desc: '描述',
      action: '编辑'
    } 
    
  ];
  handleExport() {
     
    this.href = `/api/plugin/export`;
  }
}
