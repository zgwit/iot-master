import { Component, Input } from '@angular/core';
import { RequestService } from 'src/app/request.service';
@Component({
  selector: 'app-import',
  templateUrl: './import.component.html',
  styleUrls: ['./import.component.scss'],
})
export class ImportComponent {
  @Input() url!: string ;
  uploading: Boolean = false;
  constructor(private rs: RequestService) {}
  handleImport(e: any) {
    const file: File = e.target.files[0];
    const formData = new FormData();
    formData.append('file', file);
    this.rs.post(this.url, formData).subscribe((res) => {
      console.log(res);
    });
  }
}
