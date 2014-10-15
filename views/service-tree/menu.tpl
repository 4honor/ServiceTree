{{define "menu"}}
    <div class="cont_pad">
        <div class="mt20">
            <ul class="nav nav-tabs">
                {{range $key,$val := .Menus}}
                    <li class="{{$val.Status}}"><a href="/{{$val.Name}}">{{$val.DisplayName}}</a></li>
                {{end}}
            </ul>
        </div>
    </div>
{{end}}
