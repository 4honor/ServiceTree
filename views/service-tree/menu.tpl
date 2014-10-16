{{define "menu"}}
    <div class="op_menu">
        {{range $key,$val := .Menus}}
            <a class="{{$val.Status}}" href="/{{$val.Name}}">{{$val.DisplayName}}</a>
        {{end}}
    </div>       
{{end}}
