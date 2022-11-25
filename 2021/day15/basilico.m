% from https://adventofcode.com/2021/day/15

clear all
clc

global startCell
global goalCell
global OPEN
OPEN={};
EXT=[]; % extended list

%% set M
M=[
    1 1 6 3 7 5 1 7 4 2;
    1 3 8 1 3 7 3 6 7 2;
    2 1 3 6 5 1 1 3 2 8;
    3 6 9 4 9 3 1 5 6 9;
    7 4 6 3 4 1 7 1 1 1;
    1 3 1 9 1 2 8 1 3 7;
    1 3 5 9 9 1 2 4 2 1;
    3 1 2 5 4 2 1 6 3 9;
    1 2 9 3 1 3 8 5 2 1;
    2 3 1 1 9 4 4 5 8 1;
    ];

startCell=1;
goalCell=100;

% starting node
s.g=0; % cost to come
s.h=0; % cost to go
s.path=[startCell]; % current path

insert(s);

while(~isempty(OPEN))
    ex=pop(); % expand
    if ismember(ex.path(end), EXT)
        continue % prune
    elseif isGoal(ex) % goal check
        disp("Goal found")
        disp(['Path:' num2str(ex.path)])
        disp(['g:' num2str(ex.g)])
        break
    else
        EXT(end+1)=ex.path(end);
        % expansion
        children=adj(ex.path(end),M);
        for c=1:numel(children)
            if ~ismember(children(c), ex.path) % avoid loops
                newNode.g=ex.g+M(children(c));
                newNode.h=0;
                newNode.path = [ex.path children(c)];
                insert(newNode);
            end
        end
    end
end

% given the name of a cell (linear index) returns the names of the cells
function adj = adj(i,M)
% adjacent to it
    sz=size(M);
    [r c]=ind2sub(size(M),i); % find subscripts
    adj=[r+1, c; r-1, c; r, c+1; r, c-1]; % 4-connected grid
    % remove out of bound movements
    adj=adj(~sum(adj>[sz(2),sz(1)]|adj<[1,1],2,'native'),:);
    adj=sub2ind(sz,adj(:,1),adj(:,2));
end

% goal check
function b = isGoal(n)
    global goalCell;
    if(n.path(end) == goalCell)
        b=true;
    else
        b=false;
    end
end

function el = pop()
    global OPEN;
    if(~isempty(OPEN))
        el=OPEN{1};
        OPEN=OPEN(2:end);
    else
        el=false;
    end
end

function b = push(el)
    global OPEN;
    OPEN={el OPEN{1:end}};
    b=true;
end

function i = insert(el)
    global OPEN;
    OPEN={el OPEN{1:end}}; % head insert
    for l=2:numel(OPEN) % push toward end
        if(OPEN{l-1}.g+OPEN{l-1}.h>OPEN{l}.g+OPEN{l}.h)
            swap(l-1,l);
        else
            break;
        end
    end
end

function swap(i,j)
    global OPEN;
    tmp=OPEN{i};
    OPEN{i}=OPEN{j};
    OPEN{j}=tmp;
end