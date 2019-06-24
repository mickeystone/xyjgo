// 通过读取xsd的数据获取的结构，用来读取customUI的数据

package main

// 0
type rbcustomUI struct {
	OnLoad    string `xml:"onLoad,attr"`
	LoadImage string `xml:"loadImage,attr"`

	Commands []*rbcommands `xml:"commands"`
	Ribbon   []*rbribbon   `xml:"ribbon"`
}

// 1
type rbcommands struct {
	Command []*rbcommand `xml:"command"`
}

// 2
type rbcommand struct {
	OnAction   string `xml:"onAction,attr"`
	Enabled    string `xml:"enabled,attr"`
	GetEnabled string `xml:"getEnabled,attr"`
	IdMso      string `xml:"idMso,attr"`
}

// 3
type rbribbon struct {
	StartFromScratch string `xml:"startFromScratch,attr"`

	OfficeMenu     []*rbofficeMenu     `xml:"officeMenu"`
	Qat            []*rbqat            `xml:"qat"`
	Tabs           []*rbtabs           `xml:"tabs"`
	ContextualTabs []*rbcontextualTabs `xml:"contextualTabs"`
}

// 4
type rbofficeMenu struct {
	CheckBox      []*rbcheckBox      `xml:"checkBox"`
	ToggleButton  []*rbtoggleButton  `xml:"toggleButton"`
	MenuSeparator []*rbmenuSeparator `xml:"menuSeparator"`
	SplitButton   []*rbsplitButton   `xml:"splitButton"`
	Control       []*rbcontrol       `xml:"control"`
	Button        []*rbbutton        `xml:"button"`
	Gallery       []*rbgallery       `xml:"gallery"`
	Menu          []*rbmenu          `xml:"menu"`
	DynamicMenu   []*rbdynamicMenu   `xml:"dynamicMenu"`
}

// 5
type rbcontrolBase struct {
	GetVisible      string `xml:"getVisible,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Label           string `xml:"label,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	Image           string `xml:"image,attr"`
	Screentip       string `xml:"screentip,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	GetImage        string `xml:"getImage,attr"`
	Supertip        string `xml:"supertip,attr"`
	Enabled         string `xml:"enabled,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	ShowImage       string `xml:"showImage,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	Visible         string `xml:"visible,attr"`
	Keytip          string `xml:"keytip,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
}

// 6
type rbcontrol struct {
	IdQ             string `xml:"idQ,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	Tag             string `xml:"tag,attr"`
	Id              string `xml:"id,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	Supertip        string `xml:"supertip,attr"`
	Visible         string `xml:"visible,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	Keytip          string `xml:"keytip,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	Enabled         string `xml:"enabled,attr"`
	GetImage        string `xml:"getImage,attr"`
	Label           string `xml:"label,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	ShowImage       string `xml:"showImage,attr"`
	Image           string `xml:"image,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	Screentip       string `xml:"screentip,attr"`
	IdMso           string `xml:"idMso,attr"`
}

// 7
type rbcheckBox struct {
	GetSupertip     string `xml:"getSupertip,attr"`
	Description     string `xml:"description,attr"`
	Supertip        string `xml:"supertip,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	Label           string `xml:"label,attr"`
	GetDescription  string `xml:"getDescription,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	Id              string `xml:"id,attr"`
	IdMso           string `xml:"idMso,attr"`
	Enabled         string `xml:"enabled,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Tag             string `xml:"tag,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	OnAction        string `xml:"onAction,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	Keytip          string `xml:"keytip,attr"`
	IdQ             string `xml:"idQ,attr"`
	Screentip       string `xml:"screentip,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	Visible         string `xml:"visible,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
}

// 8
type rbitem struct {
	Label     string `xml:"label,attr"`
	Image     string `xml:"image,attr"`
	ImageMso  string `xml:"imageMso,attr"`
	Screentip string `xml:"screentip,attr"`
	Supertip  string `xml:"supertip,attr"`
	Id        string `xml:"id,attr"`
}

// 9
type rbmenuSeparator struct {
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Title           string `xml:"title,attr"`
	GetTitle        string `xml:"getTitle,attr"`
	Id              string `xml:"id,attr"`
	IdQ             string `xml:"idQ,attr"`
}

// 10
type rbsplitButtonWithTitle struct {
	GetEnabled      string `xml:"getEnabled,attr"`
	Enabled         string `xml:"enabled,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	Keytip          string `xml:"keytip,attr"`
	Id              string `xml:"id,attr"`
	IdQ             string `xml:"idQ,attr"`
	Visible         string `xml:"visible,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	IdMso           string `xml:"idMso,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	Tag             string `xml:"tag,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`

	Menu         []*rbmenu         `xml:"menu"`
	Button       []*rbbutton       `xml:"button"`
	ToggleButton []*rbtoggleButton `xml:"toggleButton"`
}

// 11
type rbsplitButtonRestricted struct {
	ShowLabel       string `xml:"showLabel,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	Visible         string `xml:"visible,attr"`
	Enabled         string `xml:"enabled,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	Tag             string `xml:"tag,attr"`
	Keytip          string `xml:"keytip,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	IdQ             string `xml:"idQ,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	Id              string `xml:"id,attr"`
	IdMso           string `xml:"idMso,attr"`
}

// 12
type rbmenuWithTitle struct {
	Visible         string `xml:"visible,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	Tag             string `xml:"tag,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	ShowImage       string `xml:"showImage,attr"`
	Enabled         string `xml:"enabled,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	IdMso           string `xml:"idMso,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	Supertip        string `xml:"supertip,attr"`
	Keytip          string `xml:"keytip,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	IdQ             string `xml:"idQ,attr"`
	Screentip       string `xml:"screentip,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	Id              string `xml:"id,attr"`
	Title           string `xml:"title,attr"`
	Image           string `xml:"image,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	GetTitle        string `xml:"getTitle,attr"`
	GetImage        string `xml:"getImage,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	Label           string `xml:"label,attr"`

	Control       []*rbcontrol       `xml:"control"`
	ToggleButton  []*rbtoggleButton  `xml:"toggleButton"`
	MenuSeparator []*rbmenuSeparator `xml:"menuSeparator"`
	Button        []*rbbutton        `xml:"button"`
	CheckBox      []*rbcheckBox      `xml:"checkBox"`
	Gallery       []*rbgallery       `xml:"gallery"`
	SplitButton   []*rbsplitButton   `xml:"splitButton"`
	Menu          []*rbmenu          `xml:"menu"`
	DynamicMenu   []*rbdynamicMenu   `xml:"dynamicMenu"`
}

// 13
type rbvisibleButton struct {
	Id              string `xml:"id,attr"`
	IdMso           string `xml:"idMso,attr"`
	GetDescription  string `xml:"getDescription,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	Image           string `xml:"image,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	Supertip        string `xml:"supertip,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	IdQ             string `xml:"idQ,attr"`
	Tag             string `xml:"tag,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	Screentip       string `xml:"screentip,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	Label           string `xml:"label,attr"`
	Keytip          string `xml:"keytip,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	ShowImage       string `xml:"showImage,attr"`
	GetImage        string `xml:"getImage,attr"`
	OnAction        string `xml:"onAction,attr"`
	Description     string `xml:"description,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	Enabled         string `xml:"enabled,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
}

// 14
type rbvisibleToggleButton struct {
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	IdQ             string `xml:"idQ,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	Description     string `xml:"description,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	Id              string `xml:"id,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	ShowImage       string `xml:"showImage,attr"`
	Label           string `xml:"label,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	GetImage        string `xml:"getImage,attr"`
	Screentip       string `xml:"screentip,attr"`
	Supertip        string `xml:"supertip,attr"`
	Enabled         string `xml:"enabled,attr"`
	GetDescription  string `xml:"getDescription,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	OnAction        string `xml:"onAction,attr"`
	Image           string `xml:"image,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Keytip          string `xml:"keytip,attr"`
	IdMso           string `xml:"idMso,attr"`
	Tag             string `xml:"tag,attr"`
}

// 15
type rbqat struct {
	//	SharedControls   []*rbsharedControls   `xml:"sharedControls"`
	//	DocumentControls []*rbdocumentControls `xml:"documentControls"`
}

// 16
type rbqatItems struct {
	Control   []*rbcontrol   `xml:"control"`
	Button    []*rbbutton    `xml:"button"`
	Separator []*rbseparator `xml:"separator"`
}

// 17
type rbcontrolCloneQat struct {
	Enabled         string `xml:"enabled,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Size            string `xml:"size,attr"`
	Label           string `xml:"label,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	GetDescription  string `xml:"getDescription,attr"`
	IdMso           string `xml:"idMso,attr"`
	GetImage        string `xml:"getImage,attr"`
	Screentip       string `xml:"screentip,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	Description     string `xml:"description,attr"`
	GetSize         string `xml:"getSize,attr"`
	ShowImage       string `xml:"showImage,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	Supertip        string `xml:"supertip,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	Visible         string `xml:"visible,attr"`
	Keytip          string `xml:"keytip,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	Image           string `xml:"image,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
}

// 18
type rbseparator struct {
	Visible         string `xml:"visible,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Id              string `xml:"id,attr"`
	IdQ             string `xml:"idQ,attr"`
}

// 19
type rbtabs struct {
	Tab []*rbtab `xml:"tab"`
}

// 20
type rbtab struct {
	IdQ             string `xml:"idQ,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Visible         string `xml:"visible,attr"`
	Id              string `xml:"id,attr"`
	Label           string `xml:"label,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	IdMso           string `xml:"idMso,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	Keytip          string `xml:"keytip,attr"`
	Tag             string `xml:"tag,attr"`

	Group []*rbgroup `xml:"group"`
}

// 21
type rbgroup struct {
	Keytip          string `xml:"keytip,attr"`
	IdMso           string `xml:"idMso,attr"`
	Image           string `xml:"image,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	Screentip       string `xml:"screentip,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	Id              string `xml:"id,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	IdQ             string `xml:"idQ,attr"`
	Tag             string `xml:"tag,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	GetImage        string `xml:"getImage,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	Visible         string `xml:"visible,attr"`
	Label           string `xml:"label,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	Supertip        string `xml:"supertip,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`

	Button       []*rbbutton       `xml:"button"`
	CheckBox     []*rbcheckBox     `xml:"checkBox"`
	Control      []*rbcontrol      `xml:"control"`
	LabelControl []*rblabelControl `xml:"labelControl"`
	ComboBox     []*rbcomboBox     `xml:"comboBox"`
	SplitButton  []*rbsplitButton  `xml:"splitButton"`
	ButtonGroup  []*rbbuttonGroup  `xml:"buttonGroup"`
	Separator    []*rbseparator    `xml:"separator"`
	ToggleButton []*rbtoggleButton `xml:"toggleButton"`
	DynamicMenu  []*rbdynamicMenu  `xml:"dynamicMenu"`
	Box          []*rbbox          `xml:"box"`
	Gallery      []*rbgallery      `xml:"gallery"`
	Menu         []*rbmenu         `xml:"menu"`
	//	DropDown          []*rbdropDown          `xml:"dropDown"`
	//	DialogBoxLauncher []*rbdialogBoxLauncher `xml:"dialogBoxLauncher"`
	EditBox []*rbeditBox `xml:"editBox"`
}

// 22
type rbdialogLauncher struct {
	Button []*rbbutton `xml:"button"`
}

// 23
type rbcontrolClone struct {
	ImageMso        string `xml:"imageMso,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Description     string `xml:"description,attr"`
	GetDescription  string `xml:"getDescription,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	GetImage        string `xml:"getImage,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	Label           string `xml:"label,attr"`
	IdQ             string `xml:"idQ,attr"`
	Size            string `xml:"size,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	Visible         string `xml:"visible,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	ShowImage       string `xml:"showImage,attr"`
	Supertip        string `xml:"supertip,attr"`
	Enabled         string `xml:"enabled,attr"`
	Keytip          string `xml:"keytip,attr"`
	IdMso           string `xml:"idMso,attr"`
	GetSize         string `xml:"getSize,attr"`
	Screentip       string `xml:"screentip,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	Tag             string `xml:"tag,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	Image           string `xml:"image,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
}

// 24
type rblabelControl struct {
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Screentip       string `xml:"screentip,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	Label           string `xml:"label,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	Enabled         string `xml:"enabled,attr"`
	Tag             string `xml:"tag,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	Id              string `xml:"id,attr"`
	Visible         string `xml:"visible,attr"`
	IdQ             string `xml:"idQ,attr"`
	IdMso           string `xml:"idMso,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	Supertip        string `xml:"supertip,attr"`
}

// 25
type rbbutton struct {
	ShowLabel       string `xml:"showLabel,attr"`
	GetImage        string `xml:"getImage,attr"`
	Supertip        string `xml:"supertip,attr"`
	Label           string `xml:"label,attr"`
	Visible         string `xml:"visible,attr"`
	Size            string `xml:"size,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	IdQ             string `xml:"idQ,attr"`
	Tag             string `xml:"tag,attr"`
	GetSize         string `xml:"getSize,attr"`
	Screentip       string `xml:"screentip,attr"`
	Enabled         string `xml:"enabled,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	Keytip          string `xml:"keytip,attr"`
	Id              string `xml:"id,attr"`
	IdMso           string `xml:"idMso,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	Image           string `xml:"image,attr"`
	OnAction        string `xml:"onAction,attr"`
	Description     string `xml:"description,attr"`
	GetDescription  string `xml:"getDescription,attr"`
	ShowImage       string `xml:"showImage,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
}

// 26
type rbtoggleButton struct {
	ShowLabel       string `xml:"showLabel,attr"`
	GetImage        string `xml:"getImage,attr"`
	Keytip          string `xml:"keytip,attr"`
	Enabled         string `xml:"enabled,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	Id              string `xml:"id,attr"`
	ShowImage       string `xml:"showImage,attr"`
	IdMso           string `xml:"idMso,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	Visible         string `xml:"visible,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	Size            string `xml:"size,attr"`
	GetSize         string `xml:"getSize,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	Label           string `xml:"label,attr"`
	IdQ             string `xml:"idQ,attr"`
	Tag             string `xml:"tag,attr"`
	GetDescription  string `xml:"getDescription,attr"`
	Screentip       string `xml:"screentip,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	Supertip        string `xml:"supertip,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	Image           string `xml:"image,attr"`
	OnAction        string `xml:"onAction,attr"`
	Description     string `xml:"description,attr"`
}

// 27
type rbeditBox struct {
	GetSupertip     string `xml:"getSupertip,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	Supertip        string `xml:"supertip,attr"`
	Enabled         string `xml:"enabled,attr"`
	Label           string `xml:"label,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	Tag             string `xml:"tag,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	Screentip       string `xml:"screentip,attr"`
	IdMso           string `xml:"idMso,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	Image           string `xml:"image,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	GetImage        string `xml:"getImage,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	IdQ             string `xml:"idQ,attr"`
	ShowImage       string `xml:"showImage,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	Visible         string `xml:"visible,attr"`
	Keytip          string `xml:"keytip,attr"`
	Id              string `xml:"id,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
}

// 28
type rbcomboBox struct {
	Screentip               string `xml:"screentip,attr"`
	ShowImage               string `xml:"showImage,attr"`
	InsertBeforeQ           string `xml:"insertBeforeQ,attr"`
	GetVisible              string `xml:"getVisible,attr"`
	Tag                     string `xml:"tag,attr"`
	GetItemScreentip        string `xml:"getItemScreentip,attr"`
	GetShowLabel            string `xml:"getShowLabel,attr"`
	GetLabel                string `xml:"getLabel,attr"`
	InsertBeforeMso         string `xml:"insertBeforeMso,attr"`
	GetKeytip               string `xml:"getKeytip,attr"`
	Id                      string `xml:"id,attr"`
	IdQ                     string `xml:"idQ,attr"`
	GetItemCount            string `xml:"getItemCount,attr"`
	GetItemLabel            string `xml:"getItemLabel,attr"`
	GetItemSupertip         string `xml:"getItemSupertip,attr"`
	ShowLabel               string `xml:"showLabel,attr"`
	ImageMso                string `xml:"imageMso,attr"`
	GetScreentip            string `xml:"getScreentip,attr"`
	Enabled                 string `xml:"enabled,attr"`
	GetEnabled              string `xml:"getEnabled,attr"`
	InsertAfterMso          string `xml:"insertAfterMso,attr"`
	Visible                 string `xml:"visible,attr"`
	Keytip                  string `xml:"keytip,attr"`
	ShowItemImage           string `xml:"showItemImage,attr"`
	InvalidateContentOnDrop string `xml:"invalidateContentOnDrop,attr"`
	GetShowImage            string `xml:"getShowImage,attr"`
	Image                   string `xml:"image,attr"`
	GetSupertip             string `xml:"getSupertip,attr"`
	GetImage                string `xml:"getImage,attr"`
	IdMso                   string `xml:"idMso,attr"`
	GetItemID               string `xml:"getItemID,attr"`
	SizeString              string `xml:"sizeString,attr"`
	Supertip                string `xml:"supertip,attr"`
	Label                   string `xml:"label,attr"`
	InsertAfterQ            string `xml:"insertAfterQ,attr"`
	GetItemImage            string `xml:"getItemImage,attr"`

	Item []*rbitem `xml:"item"`
}

// 29
type rbgallery struct {
	Keytip                  string `xml:"keytip,attr"`
	GetKeytip               string `xml:"getKeytip,attr"`
	ShowItemImage           string `xml:"showItemImage,attr"`
	ImageMso                string `xml:"imageMso,attr"`
	InsertAfterMso          string `xml:"insertAfterMso,attr"`
	Supertip                string `xml:"supertip,attr"`
	GetSupertip             string `xml:"getSupertip,attr"`
	OnAction                string `xml:"onAction,attr"`
	GetItemScreentip        string `xml:"getItemScreentip,attr"`
	GetDescription          string `xml:"getDescription,attr"`
	ShowLabel               string `xml:"showLabel,attr"`
	GetShowLabel            string `xml:"getShowLabel,attr"`
	Visible                 string `xml:"visible,attr"`
	IdMso                   string `xml:"idMso,attr"`
	GetItemSupertip         string `xml:"getItemSupertip,attr"`
	Description             string `xml:"description,attr"`
	InvalidateContentOnDrop string `xml:"invalidateContentOnDrop,attr"`
	ShowImage               string `xml:"showImage,attr"`
	InsertBeforeMso         string `xml:"insertBeforeMso,attr"`
	GetVisible              string `xml:"getVisible,attr"`
	Id                      string `xml:"id,attr"`
	Tag                     string `xml:"tag,attr"`
	GetItemCount            string `xml:"getItemCount,attr"`
	GetItemID               string `xml:"getItemID,attr"`
	GetImage                string `xml:"getImage,attr"`
	GetLabel                string `xml:"getLabel,attr"`
	GetItemLabel            string `xml:"getItemLabel,attr"`
	SizeString              string `xml:"sizeString,attr"`
	GetSize                 string `xml:"getSize,attr"`
	Label                   string `xml:"label,attr"`
	InsertAfterQ            string `xml:"insertAfterQ,attr"`
	GetScreentip            string `xml:"getScreentip,attr"`
	Enabled                 string `xml:"enabled,attr"`
	IdQ                     string `xml:"idQ,attr"`
	GetItemImage            string `xml:"getItemImage,attr"`
	GetShowImage            string `xml:"getShowImage,attr"`
	Image                   string `xml:"image,attr"`
	Screentip               string `xml:"screentip,attr"`
	GetEnabled              string `xml:"getEnabled,attr"`
	InsertBeforeQ           string `xml:"insertBeforeQ,attr"`
	Size                    string `xml:"size,attr"`

	Item   []*rbitem   `xml:"item"`
	Button []*rbbutton `xml:"button"`
}

// 30
type rbmenu struct {
	Screentip       string `xml:"screentip,attr"`
	Supertip        string `xml:"supertip,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	Id              string `xml:"id,attr"`
	Enabled         string `xml:"enabled,attr"`
	GetLabel        string `xml:"getLabel,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	Label           string `xml:"label,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Image           string `xml:"image,attr"`
	GetImage        string `xml:"getImage,attr"`
	Keytip          string `xml:"keytip,attr"`
	Description     string `xml:"description,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	Visible         string `xml:"visible,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	IdMso           string `xml:"idMso,attr"`
	Size            string `xml:"size,attr"`
	GetSize         string `xml:"getSize,attr"`
	GetShowImage    string `xml:"getShowImage,attr"`
	ImageMso        string `xml:"imageMso,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	Tag             string `xml:"tag,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	GetDescription  string `xml:"getDescription,attr"`
	GetScreentip    string `xml:"getScreentip,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	IdQ             string `xml:"idQ,attr"`
	ShowImage       string `xml:"showImage,attr"`
	GetSupertip     string `xml:"getSupertip,attr"`

	Button        []*rbbutton        `xml:"button"`
	Gallery       []*rbgallery       `xml:"gallery"`
	ToggleButton  []*rbtoggleButton  `xml:"toggleButton"`
	Control       []*rbcontrol       `xml:"control"`
	CheckBox      []*rbcheckBox      `xml:"checkBox"`
	MenuSeparator []*rbmenuSeparator `xml:"menuSeparator"`
	SplitButton   []*rbsplitButton   `xml:"splitButton"`
	Menu          []*rbmenu          `xml:"menu"`
	DynamicMenu   []*rbdynamicMenu   `xml:"dynamicMenu"`
}

// 31
type rbdynamicMenu struct {
	Id                      string `xml:"id,attr"`
	ShowLabel               string `xml:"showLabel,attr"`
	ImageMso                string `xml:"imageMso,attr"`
	GetLabel                string `xml:"getLabel,attr"`
	GetContent              string `xml:"getContent,attr"`
	GetShowImage            string `xml:"getShowImage,attr"`
	Screentip               string `xml:"screentip,attr"`
	GetSupertip             string `xml:"getSupertip,attr"`
	GetEnabled              string `xml:"getEnabled,attr"`
	GetKeytip               string `xml:"getKeytip,attr"`
	InsertAfterQ            string `xml:"insertAfterQ,attr"`
	InsertBeforeQ           string `xml:"insertBeforeQ,attr"`
	InvalidateContentOnDrop string `xml:"invalidateContentOnDrop,attr"`
	GetShowLabel            string `xml:"getShowLabel,attr"`
	GetImage                string `xml:"getImage,attr"`
	Label                   string `xml:"label,attr"`
	InsertBeforeMso         string `xml:"insertBeforeMso,attr"`
	GetSize                 string `xml:"getSize,attr"`
	IdMso                   string `xml:"idMso,attr"`
	Supertip                string `xml:"supertip,attr"`
	InsertAfterMso          string `xml:"insertAfterMso,attr"`
	Visible                 string `xml:"visible,attr"`
	Description             string `xml:"description,attr"`
	IdQ                     string `xml:"idQ,attr"`
	Tag                     string `xml:"tag,attr"`
	ShowImage               string `xml:"showImage,attr"`
	Image                   string `xml:"image,attr"`
	GetVisible              string `xml:"getVisible,attr"`
	GetDescription          string `xml:"getDescription,attr"`
	Size                    string `xml:"size,attr"`
	GetScreentip            string `xml:"getScreentip,attr"`
	Enabled                 string `xml:"enabled,attr"`
	Keytip                  string `xml:"keytip,attr"`
}

// 32
type rbsplitButton struct {
	GetVisible      string `xml:"getVisible,attr"`
	Tag             string `xml:"tag,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	Visible         string `xml:"visible,attr"`
	Id              string `xml:"id,attr"`
	IdQ             string `xml:"idQ,attr"`
	GetShowLabel    string `xml:"getShowLabel,attr"`
	IdMso           string `xml:"idMso,attr"`
	GetKeytip       string `xml:"getKeytip,attr"`
	ShowLabel       string `xml:"showLabel,attr"`
	GetEnabled      string `xml:"getEnabled,attr"`
	Keytip          string `xml:"keytip,attr"`
	Enabled         string `xml:"enabled,attr"`
	Size            string `xml:"size,attr"`
	GetSize         string `xml:"getSize,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`

	Menu         []*rbmenu         `xml:"menu"`
	Button       []*rbbutton       `xml:"button"`
	ToggleButton []*rbtoggleButton `xml:"toggleButton"`
}

// 33
type rbbox struct {
	Visible         string `xml:"visible,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	BoxStyle        string `xml:"boxStyle,attr"`
	IdQ             string `xml:"idQ,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	Id              string `xml:"id,attr"`

	//	DropDown     []*rbdropDown     `xml:"dropDown"`
	DynamicMenu  []*rbdynamicMenu  `xml:"dynamicMenu"`
	ButtonGroup  []*rbbuttonGroup  `xml:"buttonGroup"`
	LabelControl []*rblabelControl `xml:"labelControl"`
	Button       []*rbbutton       `xml:"button"`
	EditBox      []*rbeditBox      `xml:"editBox"`
	Menu         []*rbmenu         `xml:"menu"`
	CheckBox     []*rbcheckBox     `xml:"checkBox"`
	ComboBox     []*rbcomboBox     `xml:"comboBox"`
	Gallery      []*rbgallery      `xml:"gallery"`
	SplitButton  []*rbsplitButton  `xml:"splitButton"`
	Box          []*rbbox          `xml:"box"`
	Control      []*rbcontrol      `xml:"control"`
	ToggleButton []*rbtoggleButton `xml:"toggleButton"`
}

// 34
type rbbuttonGroup struct {
	Visible         string `xml:"visible,attr"`
	GetVisible      string `xml:"getVisible,attr"`
	InsertAfterMso  string `xml:"insertAfterMso,attr"`
	InsertBeforeMso string `xml:"insertBeforeMso,attr"`
	InsertAfterQ    string `xml:"insertAfterQ,attr"`
	InsertBeforeQ   string `xml:"insertBeforeQ,attr"`
	Id              string `xml:"id,attr"`
	IdQ             string `xml:"idQ,attr"`

	Gallery      []*rbgallery      `xml:"gallery"`
	Menu         []*rbmenu         `xml:"menu"`
	DynamicMenu  []*rbdynamicMenu  `xml:"dynamicMenu"`
	SplitButton  []*rbsplitButton  `xml:"splitButton"`
	Control      []*rbcontrol      `xml:"control"`
	Button       []*rbbutton       `xml:"button"`
	ToggleButton []*rbtoggleButton `xml:"toggleButton"`
}

// 35
type rbcontextualTabs struct {
	TabSet []*rbtabSet `xml:"tabSet"`
}

// 36
type rbtabSet struct {
	IdMso      string `xml:"idMso,attr"`
	Visible    string `xml:"visible,attr"`
	GetVisible string `xml:"getVisible,attr"`

	Tab []*rbtab `xml:"tab"`
}
