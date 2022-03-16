import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {PageNotFoundComponent} from "../page-not-found/page-not-found.component";
import {ViewComponent} from "./view/view.component";
import {EditorComponent} from "./editor/editor.component";

const routes: Routes = [
  {path: '', component: ViewComponent},
  {path: 'view', component: ViewComponent},
  {path: 'editor', component: EditorComponent},
  {path: '**', component: PageNotFoundComponent},
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class HmiRoutingModule {
}
