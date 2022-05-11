import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {PageNotFoundComponent} from "../page-not-found/page-not-found.component";
import {EditorComponent} from "./editor/editor.component";
import {ViewerComponent} from "./viewer/viewer.component";


const routes: Routes = [
  {path: '', component: ViewerComponent},
  {path: 'viewer', component: ViewerComponent},
  {path: 'editor', component: EditorComponent},
  {path: '**', component: PageNotFoundComponent},
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class HmiRoutingModule {
}
