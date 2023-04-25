import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { isIncludeAdmin } from '../../../public';

@Component({
  selector: 'app-products-edit',
  templateUrl: './product-edit.component.html',
  styleUrls: ['./product-edit.component.scss'],
})
export class ProductEditComponent implements OnInit {
  @ViewChild('componentChild') componentChild: any;
  @Input() id: any = ''
  constructor(private router: Router, private route: ActivatedRoute) { }
  ngOnInit(): void { }
  submit() {
    this.componentChild.submit()
  }
  handleCancel() {
    const path = `${isIncludeAdmin()}/product/list`;
    this.router.navigateByUrl(path);
  }
}
